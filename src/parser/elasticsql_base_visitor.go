// Code generated from /Users/liucheng/workspace_go/elastic-sql-go/src/parser/ElasticSQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ElasticSQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseElasticSQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseElasticSQLVisitor) VisitElasticSQL(ctx *ElasticSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitCollapseExpr(ctx *CollapseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitCustomScoreExpr(ctx *CustomScoreExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitRescoreExpr(ctx *RescoreExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMemSort(ctx *MemSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitInnerHit(ctx *InnerHitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitScriptFields(ctx *ScriptFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitScriptField(ctx *ScriptFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitHighlight(ctx *HighlightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFieldAs(ctx *FieldAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitWhereExpression(ctx *WhereExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitComparableExpression(ctx *ComparableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitTermCompare(ctx *TermCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBtwCompare(ctx *BtwCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitArithmeticExpressionCompare(ctx *ArithmeticExpressionCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitArithmeticExpression(ctx *ArithmeticExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAddition(ctx *AdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMulti(ctx *MultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFunctionalCompare(ctx *FunctionalCompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitJoinFunction(ctx *JoinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitJoinFunctionNames(ctx *JoinFunctionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitScriptFunction(ctx *ScriptFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFullLevelFunction(ctx *FullLevelFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFullLevelFunctionNames(ctx *FullLevelFunctionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitTermLevelFunction(ctx *TermLevelFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitTermLevelFunctionNames(ctx *TermLevelFunctionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitProp(ctx *PropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitParam2(ctx *Param2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitExportField(ctx *ExportFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAnalysisStatement(ctx *AnalysisStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAggStatement(ctx *AggStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMetricAgg(ctx *MetricAggContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMetricNames(ctx *MetricNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMetricParams(ctx *MetricParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMkv(ctx *MkvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitMetricParamNames(ctx *MetricParamNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitScriptPhrase(ctx *ScriptPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBucketAggList(ctx *BucketAggListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBucketAgg(ctx *BucketAggContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBucketAggChoice(ctx *BucketAggChoiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitTermsBucket(ctx *TermsBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitHavingExpr(ctx *HavingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitRangeBucket(ctx *RangeBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitRangeExpr(ctx *RangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitRangeUnit(ctx *RangeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitRangeFromTo(ctx *RangeFromToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDateRangeBucket(ctx *DateRangeBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDateRangeExpr(ctx *DateRangeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDateRange(ctx *DateRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDateRangeFromTo(ctx *DateRangeFromToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitHistogramBucket(ctx *HistogramBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDateHistogramBucket(ctx *DateHistogramBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitSignificantBucket(ctx *SignificantBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFiltersBucket(ctx *FiltersBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitSparkStatement(ctx *SparkStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDataStruct(ctx *DataStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitJoinQueryAnalysisStatement(ctx *JoinQueryAnalysisStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFieldList(ctx *FieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitParamValues(ctx *ParamValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitUpdateReplaceField(ctx *UpdateReplaceFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitUpdateAddField(ctx *UpdateAddFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitUpdateRemoveField(ctx *UpdateRemoveFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBatchUpdateStatement(ctx *BatchUpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDescStatement(ctx *DescStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAddAlias(ctx *AddAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDeleteAlias(ctx *DeleteAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAlterStatement(ctx *AlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitHive2Statement(ctx *Hive2StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitJdbcStatement(ctx *JdbcStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBasicSQL(ctx *BasicSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFileLoadStatement(ctx *FileLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFieldDefine(ctx *FieldDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitHanLPStatement(ctx *HanLPStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitNlpFunc(ctx *NlpFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitPathIdentifier(ctx *PathIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitIndexIdentifier(ctx *IndexIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitAsIdentifier(ctx *AsIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitStrictIdentifier(ctx *StrictIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseElasticSQLVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}
