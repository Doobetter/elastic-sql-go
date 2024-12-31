/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import "github.com/olivere/elastic/v7"

// BtwComparableExpression
// field between a and b 默认是都包含边界
// field range [a,b] or range(a,b)
type BtwComparableExpression struct {
	ComparableExpression
	A   interface{}
	B   interface{}
	Gte bool // 是否左包含
	Lte bool // 是否右包含
}

func NewBtwComparableExpression() *BtwComparableExpression {
	expr := new(BtwComparableExpression)
	return expr
}

func (e *BtwComparableExpression) ToQueryBuilder() elastic.Query {
	queryBuilder := elastic.NewRangeQuery(e.Field)
	if e.A != nil {
		if e.Gte {
			queryBuilder.Gte(e.A)
		} else {
			queryBuilder.Gt(e.A)
		}
	}
	if e.B != nil {
		if e.Lte {
			queryBuilder.Lte(e.B)
		} else {
			queryBuilder.Lt(e.B)
		}
	}

	return queryBuilder
}
