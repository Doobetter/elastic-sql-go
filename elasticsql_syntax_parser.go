/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package elasticsql

import (
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/Doobetter/elastic-sql-go/src/parser"
	"github.com/Doobetter/elastic-sql-go/src/query"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
)

type MyElasticVisitor struct {
	parser.BaseElasticSQLVisitor
	Sql          string
	ElasticSQL   *basic.ExeElasticSQLCtx
	unitSequence int
}

func NewMyElasticVisitor(SQL string, elasticSQL *basic.ExeElasticSQLCtx) *MyElasticVisitor {
	visitor := new(MyElasticVisitor)
	visitor.Sql = SQL
	visitor.ElasticSQL = elasticSQL
	//visitor.BaseElasticSQLVisitor = new(parser.BaseElasticSQLVisitor)
	return visitor
}

func (v *MyElasticVisitor) VisitElasticSQL(ctx *parser.ElasticSQLContext) interface{} {
	if v.ElasticSQL == nil {
		v.ElasticSQL = basic.NewExeElasticSQLCtx()
		v.ElasticSQL.SQL = v.Sql
	}

	//ctx.AllStatement()
	stats := ctx.GetStatements()
	//stats := ctx.AllStatement()
	for _, stat := range stats {
		//statement := stat.Accept(v).(basic.Statement)
		statement := v.VisitStatement(stat.GetChild(0).(antlr.ParseTree)).(basic.Statement)
		v.ElasticSQL.AddStatement(statement.GetName(), statement)
		//v.VisitStatement(stat.GetChild(0).(antlr.ParseTree))

	}
	if len(v.ElasticSQL.ProcessUnitsMap) <= 0 {
		panic("解析出的STATEMENT个数为0, SQL=" + v.Sql)
	}

	return v.ElasticSQL
}

func (v *MyElasticVisitor) VisitStatement(tree antlr.ParseTree) interface{} {
	// todo other statement
	switch val := tree.(type) {
	case *parser.QueryStatementContext:
		return v.VisitQueryStatement(val)
	case *parser.AggStatementContext:
		return v.VisitAggStatement(val)
	default:
		panic("Unknown context")
	}
	return nil
}

func (v *MyElasticVisitor) VisitQueryStatement(ctx *parser.QueryStatementContext) interface{} {
	stat := query.NewQueryStatement()

	// scroll
	if ctx.GetKeep() != nil {
		stat.ScrollKeepTime = ctx.GetKeep().GetText()
	}
	if ctx.GetMinScore() != nil {
		// 通过得分过滤
		stat.MinScore, _ = strconv.ParseFloat(ctx.GetMinScore().GetText(), 64)
	}
	for _, itemI := range ctx.GetSelectItems() {
		item := itemI.(*parser.SelectItemContext)
		if item.FieldIdentifier() != nil {
			f := item.GetText()
			stat.Fields = append(stat.Fields, f)
		} else if item.Highlight() != nil {
			if stat.Highlight == nil {
				stat.Highlight = grammer.NewHighlightAdapter()
			}
			hgCtx := item.Highlight().(*parser.HighlightContext)
			var fieldSchema []string
			for _, fieldAs := range hgCtx.AllFieldAs() {
				f := fieldAs.GetField().GetText()
				as := f
				if fieldAs.GetAs() != nil {
					as = fieldAs.GetAs().GetText()
				}
				stat.Highlight.FieldAndSchema[f] = as
				fieldSchema = append(fieldSchema, as)

			}
		} else if item.ScriptField() != nil {
			if stat.ExprFields == nil {
				stat.ExprFields = make(map[string]*grammer.ScriptAdapter)
			}
			sfCtx := item.ScriptField().(*parser.ScriptFieldContext)
			as := sfCtx.GetAs().GetText()
			scriptAdapter := v.VisitScriptPhrase(sfCtx.GetScript().(*parser.ScriptPhraseContext)).(*grammer.ScriptAdapter)
			stat.ExprFields[as] = scriptAdapter
		}
	}

	for _, indexCtx := range ctx.GetIndexes() {

		indexName := v.VisitIndexName(indexCtx.GetIndex().(*parser.IndexNameContext)).(string)
		stat.Indexes = append(stat.Indexes, indexName)
		// other impl
		//stat.Indexes = append(stat.Indexes, v.VisitIndexName(indexCtx.GetChild(0).(*parser.IndexNameContext)).(string))
	}
	if ctx.GetScroll_id() != nil {
		stat.ScrollId = quotaStr(ctx.GetScroll_id().GetText())
	} else {
		// for where conditions
		if ctx.WhereExpression() != nil {
			stat.Where = v.VisitWhereExpression(ctx.WhereExpression().(*parser.WhereExpressionContext)).(grammer.Expression)
			if ctx.SCORE() != nil {
				needScore, _ := strconv.ParseBool(ctx.GetScore().GetText())
				stat.Where.SetNeedScore(needScore)
			}
			// todo custom score & rescore
		}
		if ctx.CollapseExpr() != nil {
			stat.Collapse = v.Visit(ctx.CollapseExpr()).(*grammer.CollapseAdapter)
		}

		// order by
		var sortBys []*grammer.SortAdapter
		for _, sort := range ctx.GetSorts() {
			sortBys = append(sortBys, v.VisitSortItem(sort.(*parser.SortItemContext)).(*grammer.SortAdapter))
		}
		stat.SortBys = sortBys

		// limit offset,s

		if ctx.LIMIT() != nil {
			if ctx.GetOffset() != nil {
				stat.From, _ = strconv.Atoi(ctx.GetOffset().GetText())
			}
			stat.Size, _ = strconv.Atoi(ctx.GetLimit().GetText())
		}
		if ctx.MemSort() != nil {
			// todo mem sort
		}

		if ctx.SLICE() != nil {
			stat.SliceFiled = ctx.GetSliceField().GetText()
			if ctx.GetSliceMax() != nil {
				stat.SliceMax, _ = strconv.Atoi(ctx.GetSliceMax().GetText())
			}

		}
		if n := ctx.ExportStatement(); n != nil {
			stat.Export = v.VisitExportStatement(n.(*parser.ExportStatementContext)).(*grammer.ExportClause)
			stat.Add(basic.PostProcessEnumVIAEXPORT)
		}

	}
	if ctx.GetStatName() != nil {
		mapName := ctx.GetStatName().GetText()
		stat.GenUniqName(mapName, v.unitSequence)
		v.unitSequence++
	}

	return stat
}

func quotaStr(full string) string {
	if full == "" {
		return ""
	}
	return full[1 : len(full)-1]
}
func (v *MyElasticVisitor) VisitStr(ctx *parser.StrContext) interface{} {
	if ctx.QUOTASTR() != nil {
		return quotaStr(ctx.GetText())
	}
	return ctx.GetText()
}

func (v *MyElasticVisitor) VisitIndexName(ctx *parser.IndexNameContext) interface{} {
	if ctx.QUOTASTR() != nil {
		return quotaStr(ctx.GetText())
	}
	return ctx.GetText()
}

func (v *MyElasticVisitor) VisitWhereExpression(ctx *parser.WhereExpressionContext) interface{} {
	var where grammer.Expression
	if n := ctx.LogicalExpr(); n != nil {
		where = v.VisitLogicalExpr(n.(*parser.LogicalExprContext)).(grammer.Expression)
		if ctx.SCORE() != nil {
			needScore, _ := strconv.ParseBool(ctx.GetScore().GetText())
			where.SetNeedScore(needScore)
		}
	} else {
		where = grammer.NewAllExpression()
	}
	return where
}

func (v *MyElasticVisitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	if cmp := ctx.ComparableExpression(); cmp != nil {
		return v.VisitComparableExpression(cmp.(*parser.ComparableExpressionContext))
	} else if ctx.GetInner() != nil {
		// 有括号
		expr := grammer.NewLogicExpression()
		expr.FB = true
		innerExpr := v.VisitLogicalExpr(ctx.GetInner().(*parser.LogicalExprContext))
		if _, ok := (innerExpr).(grammer.IComparableExpression); ok {
			return innerExpr
		} else if logicExpr, ok := (innerExpr).(*grammer.LogicExpression); ok {
			// Logic expression
			logicExpr.FB = true // 对于nested类型括号有特殊意义 代表两个同时满足
			return logicExpr
		} else {
			// 其他 估计也就是Logic
			return innerExpr
		}

	} else { // logic expression
		logicExpr := grammer.NewLogicExpression()
		if ctx.GetOperator().GetTokenType() == parser.ElasticSQLParserAND {
			logicExpr.Logic = grammer.LogicAND
		} else {
			logicExpr.Logic = grammer.LogicOR
		}
		leftExpr := v.VisitLogicalExpr(ctx.GetLeft().(*parser.LogicalExprContext)).(grammer.Expression)
		logicExpr.AddSubExpr(leftExpr)
		rightExpr := v.VisitLogicalExpr(ctx.GetRight().(*parser.LogicalExprContext)).(grammer.Expression)
		logicExpr.AddSubExpr(rightExpr)
		return logicExpr
	}
}

func (v *MyElasticVisitor) VisitComparableExpression(ctx *parser.ComparableExpressionContext) interface{} {
	var expr grammer.IComparableExpression
	if cmp := ctx.GetTCmp(); cmp != nil {
		expr = (v.VisitTermCompare(cmp.(*parser.TermCompareContext))).(*grammer.TermComparableExpression)
	} else if cmp := ctx.GetBtwCmp(); cmp != nil {
		expr = (v.VisitBtwCompare(cmp.(*parser.BtwCompareContext))).(grammer.IComparableExpression)
	} else if cmp := ctx.GetFuncCmp(); cmp != nil {
		expr = (v.VisitFunctionalCompare(cmp.(*parser.FunctionalCompareContext))).(grammer.IComparableExpression)
	} else if cmp := ctx.GetMathCmp(); cmp != nil {
		expr = (v.VisitArithmeticExpressionCompare(cmp.(*parser.ArithmeticExpressionCompareContext))).(grammer.IComparableExpression)
	}
	if ctx.NOT() != nil {
		// e.Not = e.Not ^ true
		if expr.GetNot() == true {
			expr.SetNot(false)
		} else {
			expr.SetNot(true)
		}
	}
	return expr
}
func (v *MyElasticVisitor) VisitTermCompare(ctx *parser.TermCompareContext) interface{} {
	expr := grammer.NewTermComparableExpression()
	field := ctx.GetField().GetText()
	expr.Paths, expr.Field = grammer.GetPathArray(field)

	operator := ctx.GetOperator().GetText()
	if "!=" == operator || "<>" == operator {
		expr.Not = true
	}
	expr.Operator = operator
	expr.Value = v.VisitParam(ctx.Param().(*parser.ParamContext))

	return expr
}

func (v *MyElasticVisitor) VisitBtwCompare(ctx *parser.BtwCompareContext) interface{} {
	expr := grammer.NewBtwComparableExpression()
	field := ctx.GetField().GetText()
	expr.Paths, expr.Field = grammer.GetPathArray(field)

	if ctx.GetGte() != nil {
		expr.Gte = true
	}
	if ctx.GetLte() != nil {
		expr.Lte = true
	}
	expr.A = v.VisitParam(ctx.GetA().(*parser.ParamContext))
	expr.B = v.VisitParam(ctx.GetB().(*parser.ParamContext))

	return expr
}

func (v *MyElasticVisitor) VisitFunctionalCompare(ctx *parser.FunctionalCompareContext) interface{} {
	if ctx.TermLevelFunction() != nil {
		expr := grammer.NewFunctionalComparableExpression()
		funcCtx := ctx.TermLevelFunction()
		expr.Func = strings.ToUpper(funcCtx.GetFuncName().GetText())
		if f := funcCtx.GetField(); f != nil {
			field := f.GetText()
			expr.Paths, expr.Field = grammer.GetPathArray(field)
		}
		if b := funcCtx.GetBoost(); b != nil {
			expr.Boost, _ = strconv.ParseFloat(b.GetText(), 64)
		}
		if params := funcCtx.GetParams(); len(params) > 0 {

			for _, p := range params {
				expr.Params = append(expr.Params, v.VisitParam2(p.(*parser.Param2Context)))
			}
		}
		if n := funcCtx.GetUseField(); n != nil {
			// todo
		}

		return expr
	} else if ctx.FullLevelFunction() != nil {
		expr := grammer.NewFullLevelFunctionalComparableExpression()
		funcCtx := ctx.FullLevelFunction()
		expr.Func = strings.ToUpper(funcCtx.GetFuncName().GetText())
		expr.Props = make(map[string]interface{})
		for _, prop := range funcCtx.GetProps() {
			expr.Props[prop.GetK().GetText()] = v.VisitParam(prop.GetV().(*parser.ParamContext))
		}
		return expr
	} else if ctx.ScriptFunction() != nil {

	} else if ctx.JoinFunction() != nil {

	}
	return nil
}

func (v *MyElasticVisitor) VisitExportStatement(ctx *parser.ExportStatementContext) interface{} {
	exportClause := grammer.NewExportClause()
	if ctx.GetFields() != nil {
		var fields []string
		for _, f := range ctx.GetFields() {
			fields = append(fields, f.GetText())
		}
		exportClause.Fields = fields
	}
	// todo distinct
	if ctx.GetFileType() != nil {
		exportClause.FileType = strings.ToUpper(ctx.GetFileType().GetText())
	}
	if exportClause.FileType == grammer.EXPORT_CSV {
		if ctx.SEP() != nil {
			exportClause.Sep = quotaStr(ctx.GetSep().GetText())
		}
		if ctx.HEADER() != nil {
			var header []string
			for _, h := range ctx.GetHeads() {
				header = append(header, v.VisitStr(h.(*parser.StrContext)).(string))
			}
			exportClause.Headers = header
		}
	}
	exportClause.FileName = v.VisitStr(ctx.GetFileName().(*parser.StrContext)).(string)

	return exportClause
}

func (v *MyElasticVisitor) VisitScriptPhrase(ctx *parser.ScriptPhraseContext) interface{} {
	// todo
	return v.VisitChildren(ctx)
}

func (v *MyElasticVisitor) VisitSortItem(ctx *parser.SortItemContext) interface{} {
	sortBy := &grammer.SortAdapter{}
	if ctx.GetField() != nil {
		sortBy.Field = ctx.GetField().GetText()
	}
	if ctx.GetScript() != nil {
		// todo
	}
	if ctx.GetOrdering() != nil {
		order := strings.ToUpper(ctx.GetOrdering().GetText())
		if "ASC" == order {
			sortBy.Ascending = true
			sortBy.OrderLogic = 1
		} else {
			sortBy.Ascending = false
			sortBy.OrderLogic = -1
		}
	}
	if ctx.GetMd() != nil {
		sortBy.Mode = ctx.GetMd().GetText()
	}
	return sortBy

}
func (v *MyElasticVisitor) VisitArrayValue(ctx *parser.ArrayValueContext) interface{} {
	return v.VisitParamValues(ctx.ParamValues().(*parser.ParamValuesContext))
}

func (v *MyElasticVisitor) VisitParam(ctx *parser.ParamContext) interface{} {
	txt := ctx.GetText()
	var value interface{}
	if ctx.LONG() != nil {
		value, _ = strconv.ParseInt(txt, 10, 64)
	} else if ctx.DOUBLE() != nil {
		value, _ = strconv.ParseFloat(txt, 64)
	} else if ctx.BooleanValue() != nil {
		value, _ = strconv.ParseBool(txt)
	} else if ctx.Str() != nil {
		value = v.VisitStr(ctx.Str().(*parser.StrContext))
	} else if ctx.ArrayValue() != nil {
		value = v.VisitArrayValue(ctx.ArrayValue().(*parser.ArrayValueContext))
	} else {
		value = txt
	}
	return value
}

func (v *MyElasticVisitor) VisitParam2(ctx *parser.Param2Context) interface{} {
	txt := ctx.GetText()
	var value interface{}
	if ctx.LONG() != nil {
		value, _ = strconv.ParseInt(txt, 10, 64)
	} else if ctx.DOUBLE() != nil {
		value, _ = strconv.ParseFloat(txt, 64)
	} else if ctx.BooleanValue() != nil {
		value, _ = strconv.ParseBool(txt)
	} else if ctx.QUOTASTR() != nil {
		value = quotaStr(ctx.QUOTASTR().GetText())
	} else if ctx.ArrayValue() != nil {
		value = v.VisitArrayValue(ctx.ArrayValue().(*parser.ArrayValueContext))
	} else {
		value = txt
	}
	return value
}

func (v *MyElasticVisitor) VisitParamValues(ctx *parser.ParamValuesContext) interface{} {
	var values []interface{}
	for _, value := range ctx.GetVs() {
		values = append(values, v.VisitParam(value.(*parser.ParamContext)))
	}
	return values
}

func (v *MyElasticVisitor) VisitAnalysisStatement(ctx *parser.AnalysisStatementContext) interface{} {

	return nil
}

// VisitAggStatement group by 部分解析
func (v *MyElasticVisitor) VisitAggStatement(ctx *parser.AggStatementContext) interface{} {

	return v.VisitChildren(ctx)
}
