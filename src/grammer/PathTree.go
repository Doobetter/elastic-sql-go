package grammer

import (
	"bytes"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/olivere/elastic/v7"
	"strings"
)

const (
	PATH_SEP_CHAR = '$'
	PATH_SEP_STR  = "$"
)

// 无nested类型的logic无PathTreee
// 为了nested类型字段的查询，同一路径下合并到一个nested查询中
type PathTree struct {
	Root        *MultiTreeNode
	NestedPaths map[string]basic.Void // 存放所有的nested path 也就是parentPath
}

func NewPathTree() *PathTree {
	tree := new(PathTree)
	tree.Root = NewMultiTreeNode("__root", true)
	return tree
}

func (t *PathTree) AddNonPathNode(expression IComparableExpression) {
	t.Root.AddData(NewTreeNode(expression))
}
func (t *PathTree) AddPathNode(expression IComparableExpression) {
	if len(expression.GetPaths()) > 0 {
		tn := NewTreeNode(expression)
		t.addChild(t.Root, tn)

	} else {
		t.AddNonPathNode(expression)
	}
}

func (t *PathTree) addChild(multiTreeNode *MultiTreeNode, node *TreeNode) {
	sub := multiTreeNode.GetChild(node.ParentPath)
	if sub != nil {
		sub.AddData(node)
	} else {
		has := false
		for path, child := range multiTreeNode.child {
			// node到当前树枝的路径是否已经存在部分，e.g. 如果树上已经有路径a, node=a.b.c的路径已经部分存在, 即已经长出了部分通向node的树杈
			if node.onPathTree(path) {
				t.addChild(child, node)
				has = true
				break
			}
		}
		// 没树枝
		if has == false {
			paths := node.Expr.GetPaths()
			mt := NewMultiTreeNode(paths[multiTreeNode.Layer], false)
			multiTreeNode.AddChild(mt)
			for i := multiTreeNode.Layer + 1; i < len(paths); i++ {
				mtn := NewMultiTreeNode(paths[i], false)
				mt.AddChild(mtn)
				mt = mtn
			}

			mt.AddData(node)
		}
	}
}
func (t *PathTree) VisitTree(tree *MultiTreeNode, boolQueryBuilder *elastic.BoolQuery, handler func(tree *MultiTreeNode, boolQuery *elastic.BoolQuery) *elastic.BoolQuery) {
	if tree == nil {
		return
	}
	bq := handler(tree, boolQueryBuilder)
	for _, child := range tree.child {
		t.VisitTree(child, bq, handler)
	}
}

// GetPathArray 工具函数 used for nested path
// e.g. "a1.b1.c1.d1" to ["a1","a1.b1","a1.b1.c1"]
// e.g "a1" 返回 空数组
func GetPathArray(field string) ([]string, string) {
	var arr []string
	fch := field
	if strings.Contains(field, PATH_SEP_STR) {
		buf := bytes.NewBufferString("")
		for _, c := range field {
			if c == PATH_SEP_CHAR {
				arr = append(arr, buf.String())
			}
			buf.WriteString(string(c))
		}
		fch = strings.ReplaceAll(fch, PATH_SEP_STR, ".")
	}
	return arr, fch
}

// MultiTreeNode 多叉树节点
type MultiTreeNode struct {
	child         map[string]*MultiTreeNode //
	Path          string
	Layer         int
	Datas         []*TreeNode
	NonRangeExprs []IComparableExpression
	RangeExprs    map[string][]*TermComparableExpression
}

func NewMultiTreeNode(path string, isRoot bool) *MultiTreeNode {
	node := new(MultiTreeNode)
	node.Path = path
	if isRoot {
		node.Layer = 0
	} else {
		node.Layer = countMatches(path, PATH_SEP_CHAR) + 1
	}
	// 采用懒惰模式 用的时候new
	//node.child = make(map[string]*MultiTreeNode)
	//node.RangeExprs = map[string]*TermComparableExpression{}
	return node
}
func (n *MultiTreeNode) AddChild(child *MultiTreeNode) {
	if n.child == nil {
		// 懒惰模式
		n.child = make(map[string]*MultiTreeNode)
	}
	n.child[child.Path] = child
}
func (n *MultiTreeNode) GetChild(path string) *MultiTreeNode {
	if n.child == nil {
		return nil
	}
	if v, ok := n.child[path]; ok {
		return v
	} else {
		return nil
	}
}

func countMatches(path string, ch int32) int {
	if path == "" {
		return 0
	}
	count := 0
	for _, c := range path {
		if c == ch {
			count++
		}
	}
	return count
}

// AddData 在节点上添加条件表达式
func (n *MultiTreeNode) AddData(node *TreeNode) {
	n.Datas = append(n.Datas, node)
	if termExpr, ok := (interface{}(node.Expr)).(*TermComparableExpression); ok {
		if termExpr.IsRange() {
			if n.RangeExprs == nil {
				n.RangeExprs = make(map[string][]*TermComparableExpression)
			}
			n.RangeExprs[termExpr.Field] = append(n.RangeExprs[termExpr.Field], termExpr)
			return
		}
	}
	n.NonRangeExprs = append(n.NonRangeExprs, node.Expr)

}

// data node 用于nested类型中存放 ComparableExpression
type TreeNode struct {
	Expr        IComparableExpression // 节点表达式
	ParentPath  string                // paths的最后一个. 如 a.b.c.d = 1 的 parentPath 是 a.b.c
	PathsOnTree []string              // paths, 如 a.b.c.d = 1 的paths is ['a','a.b','a.b.c']
	Layer       int                   // 加在第几层 len(paths)
}

func NewTreeNode(expr IComparableExpression) *TreeNode {
	node := new(TreeNode)
	node.Expr = expr
	// 是否是nested的第二层及3、4...层
	pLen := len(expr.GetPaths())
	if pLen > 0 {
		node.ParentPath = expr.GetPaths()[pLen-1]
		node.Layer = pLen
		node.PathsOnTree = append(node.PathsOnTree, expr.GetPaths()...)
	}
	return node
}

// 否在树杈path上
func (n *TreeNode) onPathTree(path string) bool {
	if n.Layer <= 0 {
		return false
	} else if n.Layer == 1 {
		return path == n.PathsOnTree[0]
	} else {
		if strings.HasPrefix(path, n.PathsOnTree[0]) {
			for i := 0; i < n.Layer; i++ {
				if n.PathsOnTree[i] == path {
					return true
				}
			}
			return false

		} else {
			return false
		}
	}
}
