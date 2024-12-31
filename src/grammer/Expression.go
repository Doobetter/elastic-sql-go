package grammer

import (
	"bytes"
	"encoding/gob"
	"github.com/olivere/elastic/v7"
	"sort"
	"strings"
)

// 编译器检测是否实现
var _ Expression = (*LogicExpression)(nil)
var _ Expression = (*ComparableExpression)(nil)

type Expression interface {
	// EnrichQueryBuilder use for single condition or '!=' or 'not' logic
	//	将自身放到boolQuery中
	EnrichBoolQueryBuilder(boolQuery *elastic.BoolQuery)
	ToQueryBuilder() elastic.Query
	GetNeedScore() bool
	//AdaptToQueryBuilder(ctx *basic.ExeElasticSQLCtx) elastic.Query
	SetNeedScore(score bool)
	IsComparableExpression() bool
	IsLogicExpression() bool
}
type ExpressionBase struct {
	NeedScore bool    // 是否算分
	Boost     float64 // for boost
}

func (e *ExpressionBase) SetNeedScore(need bool) {
	// 空实现
	e.NeedScore = need
}
func (e *ExpressionBase) GetNeedScore() bool {
	// 空实现
	return e.NeedScore
}
func (e *ExpressionBase) ToQueryBuilder() elastic.Query {
	// 空实现
	return nil
}
func (e *ExpressionBase) EnrichBoolQueryBuilder(boolQuery *elastic.BoolQuery) {
	// 空实现
}
func (e ExpressionBase) IsComparableExpression() bool {
	return false
}
func (e ExpressionBase) IsLogicExpression() bool {
	return false
}

//AdaptToQueryBuilder Expression 转化为 QueryBuilder， 在各个Statement中使用
//func (e *ExpressionBase) AdaptToQueryBuilder(ctx *basic.ExeElasticSQLCtx) elastic.Query {
//	var queryBuilder elastic.Query
//	if expr, ok := Expression(e).(*LogicExpression); ok {
//		if expr.Logic == LogicNON {
//			queryBuilder = expr.AdaptSingleComparableExpression(expr.SubExpr[0].(*ComparableExpression))
//		} else {
//			// AND OR
//			queryBuilder = e.ToQueryBuilder()
//		}
//	} else if expr, ok := Expression(e).(*ComparableExpression); ok {
//		queryBuilder = expr.AdaptSingleComparableExpression(expr)
//	} else {
//		queryBuilder = e.ToQueryBuilder()
//	}
//	if queryBuilder != nil {
//		if e.NeedScore == false {
//			queryBuilder = elastic.NewConstantScoreQuery(queryBuilder)
//		}
//	}
//	return queryBuilder
//}

// CloneNew 通过序列化深度copy.
// dst must be a pointer to the correct type
func (e *ExpressionBase) CloneNew(dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(e); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(dst)
}

// AllExpression match_all
type AllExpression struct {
	ExpressionBase
}

func NewAllExpression() *AllExpression {
	return new(AllExpression)
}
func (e *AllExpression) ToQueryBuilder() elastic.Query {
	return elastic.NewMatchAllQuery()
}

const (
	LogicAND = "AND"
	LogicOR  = "OR"
	LogicNON = "NON" //无复杂逻辑
)

// LogicExpression and or  逻辑
type LogicExpression struct {
	ExpressionBase

	SubExpr []Expression // 所有的子表达式

	PathExpr *PathTree // 用PathTree存储该logic-expr中的nested条件

	FB bool // 是否有括号强制(Force Bracket) 单独Logic 而不是和相同的path合并

	Nested        bool   // 是否有nested类型字段条件
	UniParentPath string // 子表达式的统一parentPath
	Not           bool   // 是否有 not(xxx)  专门用于 nested not
	Logic         string // 默认无LOGIC

	HasSub bool // 是否是含有子表达式 初始化后是空的为false AddSubExpr 之后就不为空 为true
}

func NewLogicExpression() *LogicExpression {
	expr := new(LogicExpression)
	expr.FB = false
	expr.Nested = false
	expr.Logic = LogicNON
	return expr
}
func (e *LogicExpression) copy(one *LogicExpression) {
	e.Nested = one.Nested
	e.SubExpr = one.SubExpr
	e.HasSub = one.HasSub
	e.PathExpr = one.PathExpr
	e.Not = one.Not
	e.Logic = one.Logic
}

func (e LogicExpression) IsLogicExpression() bool {
	return true
}

func (e *LogicExpression) AddNested() elastic.Query {
	boolQuery := elastic.NewBoolQuery()
	e.EnrichBoolQueryBuilder(boolQuery)
	// todo 判断boolQuery 是否为空boolQuery
	return boolQuery
}

func (e *LogicExpression) ToQueryBuilder() elastic.Query {
	boolQuery := elastic.NewBoolQuery()
	e.EnrichBoolQueryBuilder(boolQuery)
	// todo 判断boolQuery 是否为空boolQuery
	return boolQuery
}

// AddSubComparableExpr 将子条件类型为IComparableExpression的加入进来
func (e *LogicExpression) addSubComparableExpr(cmpExpr IComparableExpression) {

	if len(cmpExpr.GetPaths()) > 0 {
		if e.Nested == false {
			// PathTree初始化
			e.Nested = true
			e.PathExpr = NewPathTree()
		}
		e.PathExpr.AddPathNode(cmpExpr)
	} else {
		e.SubExpr = append(e.SubExpr, cmpExpr)
	}

}

// AddSubLogicExpr 加入子LogicExpr
func (e *LogicExpression) addSubLogicExpr(logicExpr *LogicExpression) {
	if logicExpr.Logic == e.Logic {
		if e.HasSub == false {
			e.copy(logicExpr)
		} else {
			if e.Nested == true && logicExpr.Nested == false {
				e.SubExpr = append(e.SubExpr, logicExpr)
			} else {
				for _, sub := range logicExpr.SubExpr {
					e.AddSubExpr(sub)
				}
			}
		}

	} else {
		// todo 同一路径下的表达式  这个比较难 且 只是特殊情况要合并 暂时不处理
		e.SubExpr = append(e.SubExpr, logicExpr)
	}

}
func (e *LogicExpression) AddSubExpr(expr Expression) {
	if cmpExpr, ok := (expr).(IComparableExpression); ok {
		e.addSubComparableExpr(cmpExpr)
	} else {
		// Logic类型
		e.addSubLogicExpr(expr.(*LogicExpression))
	}
	e.HasSub = true
}
func (e *LogicExpression) EnrichBoolQueryBuilder(boolQuery *elastic.BoolQuery) {
	if e.Logic == LogicAND {
		e.and(boolQuery)
	} else if e.Logic == LogicOR {
		e.or(boolQuery)
	} else {
		// NON 既不是AND也不是OR
		// 检测是否为nest了单个条件
		subExpr := e.SubExpr[0]
		cmpExpr, ok := (subExpr).(*ComparableExpression)
		if e.Nested && ok {
			boolQuery.Should(AdaptSingleComparableExpression(cmpExpr))
		} else {
			subExpr.EnrichBoolQueryBuilder(boolQuery)
		}
	}
}
func (e *LogicExpression) and(boolQuery *elastic.BoolQuery) {
	if e.PathExpr != nil {
		e.PathExpr.VisitTree(e.PathExpr.Root, boolQuery, func(tree *MultiTreeNode, q *elastic.BoolQuery) *elastic.BoolQuery {
			if tree.Layer == 0 {
				if len(tree.NonRangeExprs) > 0 {
					for _, expr := range tree.NonRangeExprs {
						e.genQueryFromNonRangeExprInAndLogic(expr, q)
					}
				}
				if len(tree.RangeExprs) > 0 {
					e.simpleRangeInAndLogic(q, tree.RangeExprs)
				}

				return q

			} else {
				query := elastic.NewBoolQuery()
				if len(tree.NonRangeExprs) > 0 {
					for _, expr := range tree.NonRangeExprs {
						e.genQueryFromNonRangeExprInAndLogic(expr, query)
					}
				}
				if len(tree.RangeExprs) > 0 {
					e.simpleRangeInAndLogic(query, tree.RangeExprs)
				}
				q.Must(elastic.NewNestedQuery(tree.Path, query).ScoreMode("max"))
				return query
			}
		})
	}

	for _, expr := range e.SubExpr {

		if cmpExpr, ok := (expr).(IComparableExpression); ok {
			//todo 传递context
			//expr.ctx = e.ctx
			if cmpExpr.GetNot() {
				boolQuery.MustNot(expr.ToQueryBuilder())
			} else {
				boolQuery.Must(expr.ToQueryBuilder())
			}
		} else {
			boolQuery.Must(expr.ToQueryBuilder())
		}
	}
}

func (e *LogicExpression) or(boolQuery *elastic.BoolQuery) {
	if e.PathExpr != nil {
		e.PathExpr.VisitTree(e.PathExpr.Root, boolQuery, func(tree *MultiTreeNode, q *elastic.BoolQuery) *elastic.BoolQuery {
			if tree.Layer == 0 {
				if len(tree.NonRangeExprs) > 0 {
					for _, expr := range tree.NonRangeExprs {
						e.genQueryFromNonRangeExprInOrLogic(expr, q)
					}
				}
				if len(tree.RangeExprs) > 0 {
					e.simpleRangeInOrLogic(q, tree.RangeExprs)
				}

				return q

			} else {
				query := elastic.NewBoolQuery()
				if len(tree.NonRangeExprs) > 0 {
					for _, expr := range tree.NonRangeExprs {
						e.genQueryFromNonRangeExprInOrLogic(expr, query)
					}
				}
				if len(tree.RangeExprs) > 0 {
					e.simpleRangeInOrLogic(query, tree.RangeExprs)
				}
				q.Should(elastic.NewNestedQuery(tree.Path, query).ScoreMode("max"))
				return query
			}
		})
	}
	for _, expr := range e.SubExpr {
		if cmpExpr, ok := (expr).(IComparableExpression); ok {
			//todo 传递context
			//expr.ctx = e.ctx
			if cmpExpr.GetNot() {
				subBoolQuery := elastic.NewBoolQuery()
				subBoolQuery.MustNot(expr.ToQueryBuilder())
				boolQuery.Should(subBoolQuery)
			} else {
				boolQuery.Should(expr.ToQueryBuilder())
			}
		} else {
			boolQuery.Should(expr.ToQueryBuilder())
		}
	}
}

// 对同一个field的逻辑关系处理 如 a>1 and a<4 and a<5
func (e *LogicExpression) simpleRangeInAndLogic(boolQuery *elastic.BoolQuery, commonRanges map[string][]*TermComparableExpression) {
	for _, cmps := range commonRanges {
		if len(cmps) == 1 {
			boolQuery.Filter(cmps[0].getRangeQueryBuilder())
		} else {
			sort.Sort(TermGroup(cmps))
			// 有效情况:
			// (1) a<1 and a<2 and a<3 ... 合并为a<1
			// (2) ...a>1 and a>2 and a>3 合并为a>3
			// (3)  ... and a>0 and a>1 and a<3  and a<4... 合并为a>1 and a<3
			len := len(cmps)
			st := codeForOp(cmps[0].Operator) // 1是<或者<= ，2是>或者>=
			valid := true                     // 是否有效，默认有效 出现a<0 and a>1 这个条件 条件有问题
			if st == 1 {
				// 以 < 开头
				for i := 1; i < len; i++ {
					c := codeForOp(cmps[i].Operator)
					if c == 1 {
						continue
					} else {
						valid = false // 无效
						break
					}
				}
				if valid == false {
					// 条件失效
				} else {
					//boolQuery.Must(cmps[0].getRangeQueryBuilder())
					boolQuery.Filter(cmps[0].getRangeQueryBuilder())
				}
			} else {
				// 以>开头
				validGtPos := 0 // 有效的gt下标
				validLtPos := 0
				for i := 1; i < len; i++ {
					c := codeForOp(cmps[i].Operator)
					if c == 2 {
						validGtPos = i
						if validLtPos > 0 {
							valid = false
							break
						}
					} else {
						if validLtPos == 0 {
							validLtPos = i
						}
					}
				}
				if valid == false {
					// 失效
				} else {
					q := cmps[validGtPos].getRangeQueryBuilder()
					if validLtPos > 0 {
						op := cmps[validLtPos].Operator
						v := cmps[validLtPos].Value
						if "<" == op {
							q.Lt(v)
						} else {
							q.Lte(v)
						}
					}
					//boolQuery.Must(q)
					boolQuery.Filter(q)
				}
			}

		}
	}
}

// 1是<或者<= ，2是>或者>=
func codeForOp(op string) int {
	if strings.HasPrefix(op, "<") {
		return 1
	}
	return 2
}

func (e *LogicExpression) genQueryFromNonRangeExprInAndLogic(expr IComparableExpression, boolQuery *elastic.BoolQuery) {
	//todo 传递context
	//expr.ctx = e.ctx
	termExpr, ok := (interface{}(expr)).(*TermComparableExpression)
	if ok {
		if "=" == termExpr.Operator {
			boolQuery.Filter(termExpr.ToQueryBuilder())
		} else if termExpr.Not {
			// '!=' | '<>'
			//termExpr.EnrichBoolQueryBuilder(boolQuery)
			boolQuery.MustNot(termExpr.ToQueryBuilder())
		}

	} else {
		expr.EnrichBoolQueryBuilder(boolQuery)
	}
}

func (e *LogicExpression) genQueryFromNonRangeExprInOrLogic(expr IComparableExpression, boolQuery *elastic.BoolQuery) {
	termExpr, ok := (interface{}(expr)).(*TermComparableExpression)
	if ok {
		if "=" == termExpr.Operator {
			boolQuery.Should(termExpr.ToQueryBuilder())
		} else if termExpr.Not {
			// '!=' | '<>'
			b := elastic.NewBoolQuery()
			termExpr.EnrichQueryBuilderByMustNot(b)
			boolQuery.Should(b)
		}

	} else {
		b := elastic.NewBoolQuery()
		termExpr.EnrichBoolQueryBuilder(b)
		boolQuery.Should(b)
	}
}

func (e *LogicExpression) simpleRangeInOrLogic(boolQuery *elastic.BoolQuery, commonRanges map[string][]*TermComparableExpression) {
	for _, cmps := range commonRanges {
		if len(cmps) == 1 {
			boolQuery.Should(cmps[0].getRangeQueryBuilder())
		} else {
			sort.Sort(TermGroup(cmps))
			// 有效情况:
			// (1) ...a<1 or a<2 or a<3  合并为a<3
			// (2) a>1 or a>2 or a>3 ... 合并为a>1
			// (3)  ... or a<0 or a<1 or a>3  or a>4... 合并为a<1 and a>3
			// 其他为all
			len := len(cmps)
			st := codeForOp(cmps[0].Operator) // 1是<或者<= ，2是>或者>=
			all := false                      // 是否全集 出现a>0 or a<1 这个条件 条件有问题
			if st == 2 {
				// 以 > 开头
				for i := 1; i < len; i++ {
					c := codeForOp(cmps[i].Operator)
					if c == 2 {
						continue
					} else {
						all = true
						break
					}
				}
				if all == true {
					// 全集 条件失效
				} else {
					boolQuery.Should(cmps[0].getRangeQueryBuilder())
				}
			} else {
				// 以 < 开头
				validGtPos := 0 // 有效的gt下标
				validLtPos := 0
				for i := 1; i < len; i++ {
					c := codeForOp(cmps[i].Operator)
					if c == 2 { //>

						if validLtPos == 0 {
							validGtPos = i
						} else {
							continue
						}
					} else { //<
						if validGtPos > 0 {
							//
							all = true
						}
						if validLtPos == 0 {
							validLtPos = i
						}
					}
				}
				if all {
					//  全集 条件失效
				} else {
					if validGtPos > 0 {
						boolQuery.Should(cmps[validGtPos].getRangeQueryBuilder())
					}
					boolQuery.Should(cmps[validLtPos].getRangeQueryBuilder())
				}
			}
		}
	}
}
