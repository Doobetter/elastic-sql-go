/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strconv"
	"strings"
)

var _ IComparableExpression = (*FullLevelFunctionalComparableExpression)(nil)
var _ Expression = (*FullLevelFunctionalComparableExpression)(nil)

type FullLevelFunctionalComparableExpression struct {
	ComparableExpression
	Func  string
	Props map[string]interface{}
}

func NewFullLevelFunctionalComparableExpression() *FullLevelFunctionalComparableExpression {
	expr := new(FullLevelFunctionalComparableExpression)
	expr.NeedScore = true
	return expr
}

func (e *FullLevelFunctionalComparableExpression) ToQueryBuilder() elastic.Query {

	switch e.Func {
	case FuncMATCH:
		return e.buildMatchQueryBuilder()
	case FuncMATCH_PHRASE:
		return e.buildMatchPhraseQueryBuilder()
	case FuncQUERY_STRING:
		return e.buildQueryStringQueryBuilder()
	case FuncKNN:
		return e.buildKnnQueryBuilder()

	case FuncMULTI_MATCH:
		return e.buildMultiMatchQueryBuilder()
	}
	return nil
}
func (e *FullLevelFunctionalComparableExpression) buildMatchQueryBuilder() elastic.Query {

	query := elastic.NewMatchQuery(e.Props["field"].(string), e.Props["query"])
	if analyzer, ok := e.Props["analyzer"]; ok {
		query.Analyzer(analyzer.(string))
	}
	if boost, ok := e.Props["boost"]; ok {
		// 注：本语法只有float64和int64两种数值类型
		switch boost.(type) {
		case float64:
			query.Boost(boost.(float64))
		case int64:
			query.Boost(float64(boost.(int64)))
		}

	}
	if minimumShouldMatch, ok := e.Props["minimum_should_match"]; ok {
		query.MinimumShouldMatch(minimumShouldMatch.(string))
	}
	if operator, ok := e.Props["operator"]; ok {
		query.Operator(strings.ToUpper(operator.(string)))
	} else {
		// 修改默认值为 and ，default value is OR
		query.Operator("AND")
	}

	return query
}
func (e *FullLevelFunctionalComparableExpression) buildMatchPhraseQueryBuilder() elastic.Query {
	query := elastic.NewMatchPhraseQuery(e.Props["field"].(string), e.Props["query"])
	if analyzer, ok := e.Props["analyzer"]; ok {
		query.Analyzer(analyzer.(string))
	}
	if boost, ok := e.Props["boost"]; ok {
		// 注：本语法只有float64和int64两种数值类型
		switch boost.(type) {
		case float64:
			query.Boost(boost.(float64))
		case int64:
			query.Boost(float64(boost.(int64)))
		}
	}
	if slop, ok := e.Props["slop"]; ok {
		query.Slop(int(slop.(int64)))
	}
	return query
}

// buildQueryStringQueryBuilder
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
func (e *FullLevelFunctionalComparableExpression) buildQueryStringQueryBuilder() elastic.Query {
	query := elastic.NewQueryStringQuery(e.Props["query"].(string))
	if fields, ok := e.Props["fields"]; ok {
		fs := strings.Split(fields.(string), ",")
		for _, f := range fs {
			fb := strings.Split(f, "^")
			if len(fb) == 2 {
				b, _ := strconv.ParseFloat(fb[1], 32)
				query.FieldWithBoost(fb[0], b)
			} else {
				query.Field(f)
			}
		}
	}

	if analyzer, ok := e.Props["analyzer"]; ok {
		query.Analyzer(analyzer.(string))
	}
	if boost, ok := e.Props["boost"]; ok {
		// 注：本语法只有float64和int64两种数值类型
		switch boost.(type) {
		case float64:
			query.Boost(boost.(float64))
		case int64:
			query.Boost(float64(boost.(int64)))
		}
	}
	if minimumShouldMatch, ok := e.Props["minimum_should_match"]; ok {
		query.MinimumShouldMatch(minimumShouldMatch.(string))
	}
	if operator, ok := e.Props["operator"]; ok {
		query.DefaultOperator(strings.ToUpper(operator.(string)))
	} else {
		// 修改默认值为 and ，default value is OR
		query.DefaultOperator("AND")
	}
	if tp, ok := e.Props["type"]; ok {
		query.Type(tp.(string))
		if tp.(string) == "phrase" {
			if slop, ok := e.Props["slop"]; ok {
				query.PhraseSlop(int(slop.(int64)))
			}
		}
	}

	return query
}

// buildMultiMatchQueryBuilder
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html
func (e *FullLevelFunctionalComparableExpression) buildMultiMatchQueryBuilder() elastic.Query {
	query := elastic.NewMultiMatchQuery(e.Props["query"].(string))

	fs := strings.Split(e.Props["fields"].(string), ",")
	for _, f := range fs {
		fb := strings.Split(f, "^")
		if len(fb) == 2 {
			b, _ := strconv.ParseFloat(fb[1], 32)
			query.FieldWithBoost(fb[0], b)
		} else {
			query.Field(f)
		}
	}

	if analyzer, ok := e.Props["analyzer"]; ok {
		query.Analyzer(analyzer.(string))
	}
	if boost, ok := e.Props["boost"]; ok {
		// 注：本语法只有float64和int64两种数值类型
		switch boost.(type) {
		case float64:
			query.Boost(boost.(float64))
		case int64:
			query.Boost(float64(boost.(int64)))
		}
	}
	if minimumShouldMatch, ok := e.Props["minimum_should_match"]; ok {
		query.MinimumShouldMatch(minimumShouldMatch.(string))
	}
	if operator, ok := e.Props["operator"]; ok {
		query.Operator(strings.ToUpper(operator.(string)))
	} else {
		// 修改默认值为 and ，default value is OR
		query.Operator("AND")
	}
	if tp, ok := e.Props["type"]; ok {
		query.Type(tp.(string))
	}
	if slop, ok := e.Props["slop"]; ok {
		query.Slop(int(slop.(int64)))
	}
	return query
}

// buildKnnQueryBuilder  OpenSearch KnnQuery
func (e *FullLevelFunctionalComparableExpression) buildKnnQueryBuilder() elastic.Query {
	f := e.Props["field"].(string)
	qs := map[string]interface{}{
		"knn": map[string]interface{}{
			f: map[string]interface{}{
				"vector": e.Props["vector"],
				"k":      e.Props["k"],
			},
		},
	}
	source, _ := json.Marshal(qs)

	// wrapper api 必须使用base64 encoded string， 使用RawStringQuery代替
	//return elastic.NewWrapperQuery(base64.StdEncoding.EncodeToString(source))

	return elastic.NewRawStringQuery(string(source))
}
