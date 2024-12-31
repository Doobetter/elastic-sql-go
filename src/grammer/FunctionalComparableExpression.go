/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import "github.com/olivere/elastic/v7"

var _ IComparableExpression = (*FunctionalComparableExpression)(nil)
var _ Expression = (*FunctionalComparableExpression)(nil)

type FunctionalComparableExpression struct {
	ComparableExpression
	Func     string
	Params   []interface{}
	useField string // 前置stat
}

func NewFunctionalComparableExpression() *FunctionalComparableExpression {
	expr := new(FunctionalComparableExpression)
	return expr
}

func (e *FunctionalComparableExpression) ToQueryBuilder() elastic.Query {

	switch e.Func {
	case FuncEXIST:
		return elastic.NewExistsQuery(e.Field)
	case FuncMISS:
		// e.Not = e.Not ^ true
		e.XorTrue()
		return elastic.NewExistsQuery(e.Field)
	case FuncIN:
		q := elastic.NewTermsQuery(e.Field, e.Params...)
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncOUT:
		e.XorTrue()
		q := elastic.NewTermsQuery(e.Field, e.Params...)
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncHAS_ANY:
		q := elastic.NewTermsQuery(e.Field, e.Params...)
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncHAS_ALL:
		b := elastic.NewBoolQuery()
		for _, param := range e.Params {
			b.Must(elastic.NewTermQuery(e.Field, param))
		}
		return b
	case FuncRLIKE:
		q := elastic.NewRegexpQuery(e.Field, e.Params[0].(string))
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncLIKE:
		q := elastic.NewWildcardQuery(e.Field, e.Params[0].(string))
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncSTARTS_WITH:
		e.XorTrue()
		q := elastic.NewPrefixQuery(e.Field, e.Params[0].(string))
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncIDS:
		var ids []string
		for _, param := range e.Params {
			ids = append(ids, param.(string))
		}
		q := elastic.NewIdsQuery().Ids(ids...)
		if e.Boost > 0 {
			q.Boost(e.Boost)
		}
		return q
	case FuncGEO_BBOX:
		//todo
	}

	return nil
}

const (
	FuncEXIST = "EXIST"
	FuncMISS  = "MISS"

	FuncIN  = "IN"
	FuncOUT = "OUT"

	FuncHAS_ANY = "HAS_ANY"
	FuncHAS_ALL = "HAS_ALL"

	FuncRLIKE    = "RLIKE"
	FuncLIKE     = "LIKE"
	FuncNOT_LIKE = "NOT_LIKE"

	FuncSTARTS_WITH = "STARTS_WITH"

	FuncGEO_BBOX = "GEO_BBOX"

	FuncNOT_LOGIC = "MISS|OUT|NOT_LIKE"

	FuncQUERY_STRING = "QUERY_STRING"

	FuncMULTI_MATCH = "MULTI_MATCH"

	FuncMATCH_PHRASE = "MATCH_PHRASE"

	FuncMATCH_PHRASE_PREFIX = "MATCH_PHRASE_PREFIX"
	FuncKNN                 = "KNN"
	FuncIDS                 = "IDS"

	FuncMATCH = "MATCH"

	FuncLOCAL_FILE = "LOCAL_FILE"

	FuncHAS_PARENT = "HAS_PARENT"

	FuncHAS_CHILD = "HAS_CHILD"
)
