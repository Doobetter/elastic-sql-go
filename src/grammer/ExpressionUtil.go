/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import (
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/olivere/elastic/v7"
)

// AdaptToQueryBuilder where to Query
// 条件转化总入口
func AdaptToQueryBuilder(e Expression, ctx *basic.ExeElasticSQLCtx) elastic.Query {
	var queryBuilder elastic.Query
	if expr, ok := e.(*LogicExpression); ok {
		queryBuilder = expr.ToQueryBuilder()
		// 逻辑表达式
		//if expr.Logic == LogicNON{
		//	queryBuilder = AdaptSingleComparableExpression(expr.SubExpr[0].(*ComparableExpression))
		//}else {
		//	// AND OR
		//	queryBuilder = expr.ToQueryBuilder()
		//}
	} else if expr, ok := e.(IComparableExpression); ok {
		// 单个条件表达式
		queryBuilder = AdaptSingleComparableExpression(expr)
	} else {
		queryBuilder = e.ToQueryBuilder()
	}

	//if queryBuilder!=nil{
	//	if e.GetNeedScore() == false{
	//		queryBuilder = elastic.NewConstantScoreQuery(queryBuilder)
	//	}
	//}
	return queryBuilder
}

func AdaptSingleComparableExpression(compareExpr IComparableExpression) elastic.Query {
	var query elastic.Query
	if compareExpr.GetNot() == true {
		boolQueryBuilder := elastic.NewBoolQuery()
		query = boolQueryBuilder.MustNot(compareExpr.ToQueryBuilder())
	} else {
		if compareExpr.GetNeedScore() {
			query = compareExpr.ToQueryBuilder()
		} else {
			boolQueryBuilder := elastic.NewBoolQuery().Filter(compareExpr.ToQueryBuilder())
			query = boolQueryBuilder
		}
	}
	return compareExpr.NestedSingleCondition(query)
}
