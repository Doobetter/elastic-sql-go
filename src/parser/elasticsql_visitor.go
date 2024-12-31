// Code generated from /Users/liucheng/workspace_go/elastic-sql-go/src/parser/ElasticSQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ElasticSQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ElasticSQLParser.
type ElasticSQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ElasticSQLParser#elasticSQL.
	VisitElasticSQL(ctx *ElasticSQLContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#queryStatement.
	VisitQueryStatement(ctx *QueryStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#selectItem.
	VisitSelectItem(ctx *SelectItemContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#collapseExpr.
	VisitCollapseExpr(ctx *CollapseExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#customScoreExpr.
	VisitCustomScoreExpr(ctx *CustomScoreExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#rescoreExpr.
	VisitRescoreExpr(ctx *RescoreExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#memSort.
	VisitMemSort(ctx *MemSortContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#innerHit.
	VisitInnerHit(ctx *InnerHitContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#scriptField.
	VisitScriptField(ctx *ScriptFieldContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#highlight.
	VisitHighlight(ctx *HighlightContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fieldAs.
	VisitFieldAs(ctx *FieldAsContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#whereExpression.
	VisitWhereExpression(ctx *WhereExpressionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#logicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#comparableExpression.
	VisitComparableExpression(ctx *ComparableExpressionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#termCompare.
	VisitTermCompare(ctx *TermCompareContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#btwCompare.
	VisitBtwCompare(ctx *BtwCompareContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#arithmeticExpressionCompare.
	VisitArithmeticExpressionCompare(ctx *ArithmeticExpressionCompareContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#arithmeticExpression.
	VisitArithmeticExpression(ctx *ArithmeticExpressionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#multiplyingExpression.
	VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#multi.
	VisitMulti(ctx *MultiContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#functionalCompare.
	VisitFunctionalCompare(ctx *FunctionalCompareContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#joinFunction.
	VisitJoinFunction(ctx *JoinFunctionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#joinFunctionNames.
	VisitJoinFunctionNames(ctx *JoinFunctionNamesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#scriptFunction.
	VisitScriptFunction(ctx *ScriptFunctionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fullLevelFunction.
	VisitFullLevelFunction(ctx *FullLevelFunctionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fullLevelFunctionNames.
	VisitFullLevelFunctionNames(ctx *FullLevelFunctionNamesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#termLevelFunction.
	VisitTermLevelFunction(ctx *TermLevelFunctionContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#termLevelFunctionNames.
	VisitTermLevelFunctionNames(ctx *TermLevelFunctionNamesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#prop.
	VisitProp(ctx *PropContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#param2.
	VisitParam2(ctx *Param2Context) interface{}

	// Visit a parse tree produced by ElasticSQLParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#exportField.
	VisitExportField(ctx *ExportFieldContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#analysisStatement.
	VisitAnalysisStatement(ctx *AnalysisStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#aggStatement.
	VisitAggStatement(ctx *AggStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#metricAgg.
	VisitMetricAgg(ctx *MetricAggContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#metricNames.
	VisitMetricNames(ctx *MetricNamesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#metricParams.
	VisitMetricParams(ctx *MetricParamsContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#mkv.
	VisitMkv(ctx *MkvContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#metricParamNames.
	VisitMetricParamNames(ctx *MetricParamNamesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#scriptPhrase.
	VisitScriptPhrase(ctx *ScriptPhraseContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#bucketAggList.
	VisitBucketAggList(ctx *BucketAggListContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#bucketAgg.
	VisitBucketAgg(ctx *BucketAggContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#bucketAggChoice.
	VisitBucketAggChoice(ctx *BucketAggChoiceContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#termsBucket.
	VisitTermsBucket(ctx *TermsBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#havingExpr.
	VisitHavingExpr(ctx *HavingExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#rangeBucket.
	VisitRangeBucket(ctx *RangeBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#rangeExpr.
	VisitRangeExpr(ctx *RangeExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#rangeUnit.
	VisitRangeUnit(ctx *RangeUnitContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#rangeFromTo.
	VisitRangeFromTo(ctx *RangeFromToContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dateRangeBucket.
	VisitDateRangeBucket(ctx *DateRangeBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dateRangeExpr.
	VisitDateRangeExpr(ctx *DateRangeExprContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dateRange.
	VisitDateRange(ctx *DateRangeContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dateRangeFromTo.
	VisitDateRangeFromTo(ctx *DateRangeFromToContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#histogramBucket.
	VisitHistogramBucket(ctx *HistogramBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dateHistogramBucket.
	VisitDateHistogramBucket(ctx *DateHistogramBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#significantBucket.
	VisitSignificantBucket(ctx *SignificantBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#filtersBucket.
	VisitFiltersBucket(ctx *FiltersBucketContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#sparkStatement.
	VisitSparkStatement(ctx *SparkStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dataStruct.
	VisitDataStruct(ctx *DataStructContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#joinQueryAnalysisStatement.
	VisitJoinQueryAnalysisStatement(ctx *JoinQueryAnalysisStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#paramValues.
	VisitParamValues(ctx *ParamValuesContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#updateReplaceField.
	VisitUpdateReplaceField(ctx *UpdateReplaceFieldContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#updateAddField.
	VisitUpdateAddField(ctx *UpdateAddFieldContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#updateRemoveField.
	VisitUpdateRemoveField(ctx *UpdateRemoveFieldContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#batchUpdateStatement.
	VisitBatchUpdateStatement(ctx *BatchUpdateStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#descStatement.
	VisitDescStatement(ctx *DescStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#addAlias.
	VisitAddAlias(ctx *AddAliasContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#deleteAlias.
	VisitDeleteAlias(ctx *DeleteAliasContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#hive2Statement.
	VisitHive2Statement(ctx *Hive2StatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#jdbcStatement.
	VisitJdbcStatement(ctx *JdbcStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#basicSQL.
	VisitBasicSQL(ctx *BasicSQLContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fileLoadStatement.
	VisitFileLoadStatement(ctx *FileLoadStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fieldDefine.
	VisitFieldDefine(ctx *FieldDefineContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#analyzeStatement.
	VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#hanLPStatement.
	VisitHanLPStatement(ctx *HanLPStatementContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#nlpFunc.
	VisitNlpFunc(ctx *NlpFuncContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#pathIdentifier.
	VisitPathIdentifier(ctx *PathIdentifierContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#fieldIdentifier.
	VisitFieldIdentifier(ctx *FieldIdentifierContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#indexIdentifier.
	VisitIndexIdentifier(ctx *IndexIdentifierContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#strictIdentifier.
	VisitStrictIdentifier(ctx *StrictIdentifierContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by ElasticSQLParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}
}
