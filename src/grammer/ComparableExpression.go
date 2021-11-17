package grammer

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"regexp"
	"strings"
)

var _ IComparableExpression = (*ComparableExpression)(nil)

type IComparableExpression interface {
	Expression
	GetField() string
	SetField(field string)
	GetPaths() []string
	SetPaths(paths []string)
	SetNot(not bool)
	GetNot() bool
	NestedSingleCondition(query elastic.Query) elastic.Query
}


type ComparableExpression struct {
	ExpressionBase
	Field string
	Paths []string // nested field's path, e.g. 字段a.b.c.d的paths是 [a, a.b, a.b.c]
	Not   bool     // 是否是not语义, 按boolean异或生成, not not = false, not != 为 false , not = 为 true
	//NeedScore bool // 不需要score的就用filter
}

//AdaptSingleComparableExpression 单个ComparableExpression适配成QueryBuilder
//func (e *ComparableExpression) AdaptSingleComparableExpression(compareExpr *ComparableExpression) elastic.Query {
//	builder := elastic.NewBoolQuery()
//
//	compareExpr.EnrichBoolQueryBuilder(builder)
//
//	return compareExpr.NestedSingleCondition(builder)
//}

//XorTrue e.Not^true
func (e ComparableExpression) XorTrue(){
	if e.Not == true {e.Not = false}else {e.Not = true}
}


func (e ComparableExpression) GetField() string {
	return e.Field
}

func (e *ComparableExpression) SetField(field string) {
	e.Field = field
}

func (e ComparableExpression) GetPaths() []string {
	return e.Paths
}

func (e *ComparableExpression) SetPaths(paths []string) {
	e.Paths = paths
}

func (e *ComparableExpression) SetNot(not bool) {
	e.Not = not
}

func (e *ComparableExpression) GetNot() bool {
	return e.Not
}

func (e ComparableExpression) IsComparableExpression() bool {
	return true
}

func NestedSingleCondition(e *ComparableExpression, query elastic.Query) elastic.Query {
	if len(e.Paths) > 0 {
		var tmp elastic.Query
		for i := len(e.Paths) - 1; i >= 0; i-- {
			path := e.Paths[i]
			tmp = elastic.NewNestedQuery(path, query)
		}
		return tmp
	}
	return query
}
// NestedSingleCondition 单个条件时对嵌套的判断
func (e *ComparableExpression) NestedSingleCondition(query elastic.Query) elastic.Query {
	return NestedSingleCondition(e,query)
}

// EnrichQueryBuilder ComparableExpression类查询通用的对not做处理
func (e *ComparableExpression) EnrichBoolQueryBuilder(boolQueryBuilder *elastic.BoolQuery) {
	if e.Not == true {
		// boolQueryBuilder.mustNot(QueryBuilders.termQuery(field, Value));
		boolQueryBuilder.MustNot(e.ToQueryBuilder())
	} else {
		if e.NeedScore {
			boolQueryBuilder.Must(e.ToQueryBuilder())
		} else {
			boolQueryBuilder.Filter(e.ToQueryBuilder())
		}
	}
}

type TermComparableExpression struct {
	ComparableExpression
	Operator string
	Value    interface{}
}

func NewTermComparableExpression() *TermComparableExpression {
	expr := new(TermComparableExpression)
	return expr
}

var reg = regexp.MustCompile(`"==|=|!=|<>`)

func (e *TermComparableExpression) IsRange() bool {
	if reg.MatchString(e.Operator) {
		return false
	}
	return true

}

func (e *TermComparableExpression) ToQueryBuilder() elastic.Query {
	var queryBuilder elastic.Query
	if reg.MatchString(e.Operator) {
		tq := elastic.NewTermQuery(e.Field, e.Value)
		if e.Boost > 0 {
			tq.Boost(e.Boost)
		}
		queryBuilder = tq
	} else {
		queryBuilder = e.getRangeQueryBuilder()
	}
	return queryBuilder
}

func (e *TermComparableExpression) getRangeQueryBuilder() *elastic.RangeQuery {
	var rangeQuery *elastic.RangeQuery
	if ">=" == e.Operator {
		rangeQuery = elastic.NewRangeQuery(e.Field).Gte(e.Value)
	} else if ">" == e.Operator {
		rangeQuery = elastic.NewRangeQuery(e.Field).Gte(e.Value)
	} else if "<" == e.Operator {
		rangeQuery = elastic.NewRangeQuery(e.Field).Lt(e.Value)
	} else if "<=" == e.Operator {
		rangeQuery = elastic.NewRangeQuery(e.Field).Lte(e.Value)
	} else {
		panic("非法比较符" + e.Operator)
	}
	if rangeQuery != nil && e.Boost > 0 {
		rangeQuery.Boost(e.Boost)
	}
	return rangeQuery
}

func (e *TermComparableExpression) EnrichQueryBuilderByMustNot(boolQueryBuilder *elastic.BoolQuery) {
	boolQueryBuilder.MustNot(elastic.NewTermQuery(e.Field, e.Value))
}

type TermGroup []*TermComparableExpression

func (g TermGroup) Len() int {
	return len(g)
}

//Less 有小到大排序
// 本体系中只有int64 float64两种类型
func (g TermGroup) Less(i, j int) bool {
	one := g[i]
	two := g[j]
	if oneStr, ok := one.Value.(string); ok {
		twoStr := fmt.Sprintf("%v", two.Value)
		return strings.Compare(oneStr, twoStr) < 0
	} else if twoStr, ok := two.Value.(string); ok {
		oneStr := fmt.Sprintf("%v", one.Value)
		return strings.Compare(oneStr, twoStr) < 0
	}
	// number 类型
	var oneNumber float64
	switch one.Value.(type) {
	case float64:
		oneNumber = one.Value.(float64)
	case int64:
		oneNumber = float64(one.Value.(int64))
	}

	var twoNumber float64
	switch two.Value.(type) {
	case float64:
		twoNumber = two.Value.(float64)
	case int64:
		twoNumber = float64(two.Value.(int64))
	}
	rz := oneNumber - twoNumber

	return rz < 0
}

func (g TermGroup) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

