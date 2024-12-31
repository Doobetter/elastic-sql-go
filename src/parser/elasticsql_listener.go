// Code generated from /Users/liucheng/workspace_go/elastic-sql-go/src/parser/ElasticSQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ElasticSQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ElasticSQLListener is a complete listener for a parse tree produced by ElasticSQLParser.
type ElasticSQLListener interface {
	antlr.ParseTreeListener

	// EnterElasticSQL is called when entering the elasticSQL production.
	EnterElasticSQL(c *ElasticSQLContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterQueryStatement is called when entering the queryStatement production.
	EnterQueryStatement(c *QueryStatementContext)

	// EnterSelectItem is called when entering the selectItem production.
	EnterSelectItem(c *SelectItemContext)

	// EnterCollapseExpr is called when entering the collapseExpr production.
	EnterCollapseExpr(c *CollapseExprContext)

	// EnterCustomScoreExpr is called when entering the customScoreExpr production.
	EnterCustomScoreExpr(c *CustomScoreExprContext)

	// EnterRescoreExpr is called when entering the rescoreExpr production.
	EnterRescoreExpr(c *RescoreExprContext)

	// EnterMemSort is called when entering the memSort production.
	EnterMemSort(c *MemSortContext)

	// EnterInnerHit is called when entering the innerHit production.
	EnterInnerHit(c *InnerHitContext)

	// EnterScriptField is called when entering the scriptField production.
	EnterScriptField(c *ScriptFieldContext)

	// EnterHighlight is called when entering the highlight production.
	EnterHighlight(c *HighlightContext)

	// EnterFieldAs is called when entering the fieldAs production.
	EnterFieldAs(c *FieldAsContext)

	// EnterWhereExpression is called when entering the whereExpression production.
	EnterWhereExpression(c *WhereExpressionContext)

	// EnterLogicalExpr is called when entering the logicalExpr production.
	EnterLogicalExpr(c *LogicalExprContext)

	// EnterComparableExpression is called when entering the comparableExpression production.
	EnterComparableExpression(c *ComparableExpressionContext)

	// EnterTermCompare is called when entering the termCompare production.
	EnterTermCompare(c *TermCompareContext)

	// EnterBtwCompare is called when entering the btwCompare production.
	EnterBtwCompare(c *BtwCompareContext)

	// EnterArithmeticExpressionCompare is called when entering the arithmeticExpressionCompare production.
	EnterArithmeticExpressionCompare(c *ArithmeticExpressionCompareContext)

	// EnterArithmeticExpression is called when entering the arithmeticExpression production.
	EnterArithmeticExpression(c *ArithmeticExpressionContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterMulti is called when entering the multi production.
	EnterMulti(c *MultiContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterFunctionalCompare is called when entering the functionalCompare production.
	EnterFunctionalCompare(c *FunctionalCompareContext)

	// EnterJoinFunction is called when entering the joinFunction production.
	EnterJoinFunction(c *JoinFunctionContext)

	// EnterJoinFunctionNames is called when entering the joinFunctionNames production.
	EnterJoinFunctionNames(c *JoinFunctionNamesContext)

	// EnterScriptFunction is called when entering the scriptFunction production.
	EnterScriptFunction(c *ScriptFunctionContext)

	// EnterFullLevelFunction is called when entering the fullLevelFunction production.
	EnterFullLevelFunction(c *FullLevelFunctionContext)

	// EnterFullLevelFunctionNames is called when entering the fullLevelFunctionNames production.
	EnterFullLevelFunctionNames(c *FullLevelFunctionNamesContext)

	// EnterTermLevelFunction is called when entering the termLevelFunction production.
	EnterTermLevelFunction(c *TermLevelFunctionContext)

	// EnterTermLevelFunctionNames is called when entering the termLevelFunctionNames production.
	EnterTermLevelFunctionNames(c *TermLevelFunctionNamesContext)

	// EnterProp is called when entering the prop production.
	EnterProp(c *PropContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParam2 is called when entering the param2 production.
	EnterParam2(c *Param2Context)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

	// EnterExportField is called when entering the exportField production.
	EnterExportField(c *ExportFieldContext)

	// EnterAnalysisStatement is called when entering the analysisStatement production.
	EnterAnalysisStatement(c *AnalysisStatementContext)

	// EnterAggStatement is called when entering the aggStatement production.
	EnterAggStatement(c *AggStatementContext)

	// EnterMetricAgg is called when entering the metricAgg production.
	EnterMetricAgg(c *MetricAggContext)

	// EnterMetricNames is called when entering the metricNames production.
	EnterMetricNames(c *MetricNamesContext)

	// EnterMetricParams is called when entering the metricParams production.
	EnterMetricParams(c *MetricParamsContext)

	// EnterMkv is called when entering the mkv production.
	EnterMkv(c *MkvContext)

	// EnterMetricParamNames is called when entering the metricParamNames production.
	EnterMetricParamNames(c *MetricParamNamesContext)

	// EnterScriptPhrase is called when entering the scriptPhrase production.
	EnterScriptPhrase(c *ScriptPhraseContext)

	// EnterBucketAggList is called when entering the bucketAggList production.
	EnterBucketAggList(c *BucketAggListContext)

	// EnterBucketAgg is called when entering the bucketAgg production.
	EnterBucketAgg(c *BucketAggContext)

	// EnterBucketAggChoice is called when entering the bucketAggChoice production.
	EnterBucketAggChoice(c *BucketAggChoiceContext)

	// EnterTermsBucket is called when entering the termsBucket production.
	EnterTermsBucket(c *TermsBucketContext)

	// EnterHavingExpr is called when entering the havingExpr production.
	EnterHavingExpr(c *HavingExprContext)

	// EnterRangeBucket is called when entering the rangeBucket production.
	EnterRangeBucket(c *RangeBucketContext)

	// EnterRangeExpr is called when entering the rangeExpr production.
	EnterRangeExpr(c *RangeExprContext)

	// EnterRangeUnit is called when entering the rangeUnit production.
	EnterRangeUnit(c *RangeUnitContext)

	// EnterRangeFromTo is called when entering the rangeFromTo production.
	EnterRangeFromTo(c *RangeFromToContext)

	// EnterDateRangeBucket is called when entering the dateRangeBucket production.
	EnterDateRangeBucket(c *DateRangeBucketContext)

	// EnterDateRangeExpr is called when entering the dateRangeExpr production.
	EnterDateRangeExpr(c *DateRangeExprContext)

	// EnterDateRange is called when entering the dateRange production.
	EnterDateRange(c *DateRangeContext)

	// EnterDateRangeFromTo is called when entering the dateRangeFromTo production.
	EnterDateRangeFromTo(c *DateRangeFromToContext)

	// EnterHistogramBucket is called when entering the histogramBucket production.
	EnterHistogramBucket(c *HistogramBucketContext)

	// EnterDateHistogramBucket is called when entering the dateHistogramBucket production.
	EnterDateHistogramBucket(c *DateHistogramBucketContext)

	// EnterSignificantBucket is called when entering the significantBucket production.
	EnterSignificantBucket(c *SignificantBucketContext)

	// EnterFiltersBucket is called when entering the filtersBucket production.
	EnterFiltersBucket(c *FiltersBucketContext)

	// EnterSparkStatement is called when entering the sparkStatement production.
	EnterSparkStatement(c *SparkStatementContext)

	// EnterDataStruct is called when entering the dataStruct production.
	EnterDataStruct(c *DataStructContext)

	// EnterJoinQueryAnalysisStatement is called when entering the joinQueryAnalysisStatement production.
	EnterJoinQueryAnalysisStatement(c *JoinQueryAnalysisStatementContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterParamValues is called when entering the paramValues production.
	EnterParamValues(c *ParamValuesContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterUpdateReplaceField is called when entering the updateReplaceField production.
	EnterUpdateReplaceField(c *UpdateReplaceFieldContext)

	// EnterUpdateAddField is called when entering the updateAddField production.
	EnterUpdateAddField(c *UpdateAddFieldContext)

	// EnterUpdateRemoveField is called when entering the updateRemoveField production.
	EnterUpdateRemoveField(c *UpdateRemoveFieldContext)

	// EnterBatchUpdateStatement is called when entering the batchUpdateStatement production.
	EnterBatchUpdateStatement(c *BatchUpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterDescStatement is called when entering the descStatement production.
	EnterDescStatement(c *DescStatementContext)

	// EnterAddAlias is called when entering the addAlias production.
	EnterAddAlias(c *AddAliasContext)

	// EnterDeleteAlias is called when entering the deleteAlias production.
	EnterDeleteAlias(c *DeleteAliasContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterHive2Statement is called when entering the hive2Statement production.
	EnterHive2Statement(c *Hive2StatementContext)

	// EnterJdbcStatement is called when entering the jdbcStatement production.
	EnterJdbcStatement(c *JdbcStatementContext)

	// EnterBasicSQL is called when entering the basicSQL production.
	EnterBasicSQL(c *BasicSQLContext)

	// EnterFileLoadStatement is called when entering the fileLoadStatement production.
	EnterFileLoadStatement(c *FileLoadStatementContext)

	// EnterFieldDefine is called when entering the fieldDefine production.
	EnterFieldDefine(c *FieldDefineContext)

	// EnterAnalyzeStatement is called when entering the analyzeStatement production.
	EnterAnalyzeStatement(c *AnalyzeStatementContext)

	// EnterHanLPStatement is called when entering the hanLPStatement production.
	EnterHanLPStatement(c *HanLPStatementContext)

	// EnterNlpFunc is called when entering the nlpFunc production.
	EnterNlpFunc(c *NlpFuncContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterPathIdentifier is called when entering the pathIdentifier production.
	EnterPathIdentifier(c *PathIdentifierContext)

	// EnterFieldIdentifier is called when entering the fieldIdentifier production.
	EnterFieldIdentifier(c *FieldIdentifierContext)

	// EnterIndexIdentifier is called when entering the indexIdentifier production.
	EnterIndexIdentifier(c *IndexIdentifierContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterStrictIdentifier is called when entering the strictIdentifier production.
	EnterStrictIdentifier(c *StrictIdentifierContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterNonReserved is called when entering the nonReserved production.
	EnterNonReserved(c *NonReservedContext)

	// ExitElasticSQL is called when exiting the elasticSQL production.
	ExitElasticSQL(c *ElasticSQLContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitQueryStatement is called when exiting the queryStatement production.
	ExitQueryStatement(c *QueryStatementContext)

	// ExitSelectItem is called when exiting the selectItem production.
	ExitSelectItem(c *SelectItemContext)

	// ExitCollapseExpr is called when exiting the collapseExpr production.
	ExitCollapseExpr(c *CollapseExprContext)

	// ExitCustomScoreExpr is called when exiting the customScoreExpr production.
	ExitCustomScoreExpr(c *CustomScoreExprContext)

	// ExitRescoreExpr is called when exiting the rescoreExpr production.
	ExitRescoreExpr(c *RescoreExprContext)

	// ExitMemSort is called when exiting the memSort production.
	ExitMemSort(c *MemSortContext)

	// ExitInnerHit is called when exiting the innerHit production.
	ExitInnerHit(c *InnerHitContext)

	// ExitScriptField is called when exiting the scriptField production.
	ExitScriptField(c *ScriptFieldContext)

	// ExitHighlight is called when exiting the highlight production.
	ExitHighlight(c *HighlightContext)

	// ExitFieldAs is called when exiting the fieldAs production.
	ExitFieldAs(c *FieldAsContext)

	// ExitWhereExpression is called when exiting the whereExpression production.
	ExitWhereExpression(c *WhereExpressionContext)

	// ExitLogicalExpr is called when exiting the logicalExpr production.
	ExitLogicalExpr(c *LogicalExprContext)

	// ExitComparableExpression is called when exiting the comparableExpression production.
	ExitComparableExpression(c *ComparableExpressionContext)

	// ExitTermCompare is called when exiting the termCompare production.
	ExitTermCompare(c *TermCompareContext)

	// ExitBtwCompare is called when exiting the btwCompare production.
	ExitBtwCompare(c *BtwCompareContext)

	// ExitArithmeticExpressionCompare is called when exiting the arithmeticExpressionCompare production.
	ExitArithmeticExpressionCompare(c *ArithmeticExpressionCompareContext)

	// ExitArithmeticExpression is called when exiting the arithmeticExpression production.
	ExitArithmeticExpression(c *ArithmeticExpressionContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitMulti is called when exiting the multi production.
	ExitMulti(c *MultiContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitFunctionalCompare is called when exiting the functionalCompare production.
	ExitFunctionalCompare(c *FunctionalCompareContext)

	// ExitJoinFunction is called when exiting the joinFunction production.
	ExitJoinFunction(c *JoinFunctionContext)

	// ExitJoinFunctionNames is called when exiting the joinFunctionNames production.
	ExitJoinFunctionNames(c *JoinFunctionNamesContext)

	// ExitScriptFunction is called when exiting the scriptFunction production.
	ExitScriptFunction(c *ScriptFunctionContext)

	// ExitFullLevelFunction is called when exiting the fullLevelFunction production.
	ExitFullLevelFunction(c *FullLevelFunctionContext)

	// ExitFullLevelFunctionNames is called when exiting the fullLevelFunctionNames production.
	ExitFullLevelFunctionNames(c *FullLevelFunctionNamesContext)

	// ExitTermLevelFunction is called when exiting the termLevelFunction production.
	ExitTermLevelFunction(c *TermLevelFunctionContext)

	// ExitTermLevelFunctionNames is called when exiting the termLevelFunctionNames production.
	ExitTermLevelFunctionNames(c *TermLevelFunctionNamesContext)

	// ExitProp is called when exiting the prop production.
	ExitProp(c *PropContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParam2 is called when exiting the param2 production.
	ExitParam2(c *Param2Context)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

	// ExitExportField is called when exiting the exportField production.
	ExitExportField(c *ExportFieldContext)

	// ExitAnalysisStatement is called when exiting the analysisStatement production.
	ExitAnalysisStatement(c *AnalysisStatementContext)

	// ExitAggStatement is called when exiting the aggStatement production.
	ExitAggStatement(c *AggStatementContext)

	// ExitMetricAgg is called when exiting the metricAgg production.
	ExitMetricAgg(c *MetricAggContext)

	// ExitMetricNames is called when exiting the metricNames production.
	ExitMetricNames(c *MetricNamesContext)

	// ExitMetricParams is called when exiting the metricParams production.
	ExitMetricParams(c *MetricParamsContext)

	// ExitMkv is called when exiting the mkv production.
	ExitMkv(c *MkvContext)

	// ExitMetricParamNames is called when exiting the metricParamNames production.
	ExitMetricParamNames(c *MetricParamNamesContext)

	// ExitScriptPhrase is called when exiting the scriptPhrase production.
	ExitScriptPhrase(c *ScriptPhraseContext)

	// ExitBucketAggList is called when exiting the bucketAggList production.
	ExitBucketAggList(c *BucketAggListContext)

	// ExitBucketAgg is called when exiting the bucketAgg production.
	ExitBucketAgg(c *BucketAggContext)

	// ExitBucketAggChoice is called when exiting the bucketAggChoice production.
	ExitBucketAggChoice(c *BucketAggChoiceContext)

	// ExitTermsBucket is called when exiting the termsBucket production.
	ExitTermsBucket(c *TermsBucketContext)

	// ExitHavingExpr is called when exiting the havingExpr production.
	ExitHavingExpr(c *HavingExprContext)

	// ExitRangeBucket is called when exiting the rangeBucket production.
	ExitRangeBucket(c *RangeBucketContext)

	// ExitRangeExpr is called when exiting the rangeExpr production.
	ExitRangeExpr(c *RangeExprContext)

	// ExitRangeUnit is called when exiting the rangeUnit production.
	ExitRangeUnit(c *RangeUnitContext)

	// ExitRangeFromTo is called when exiting the rangeFromTo production.
	ExitRangeFromTo(c *RangeFromToContext)

	// ExitDateRangeBucket is called when exiting the dateRangeBucket production.
	ExitDateRangeBucket(c *DateRangeBucketContext)

	// ExitDateRangeExpr is called when exiting the dateRangeExpr production.
	ExitDateRangeExpr(c *DateRangeExprContext)

	// ExitDateRange is called when exiting the dateRange production.
	ExitDateRange(c *DateRangeContext)

	// ExitDateRangeFromTo is called when exiting the dateRangeFromTo production.
	ExitDateRangeFromTo(c *DateRangeFromToContext)

	// ExitHistogramBucket is called when exiting the histogramBucket production.
	ExitHistogramBucket(c *HistogramBucketContext)

	// ExitDateHistogramBucket is called when exiting the dateHistogramBucket production.
	ExitDateHistogramBucket(c *DateHistogramBucketContext)

	// ExitSignificantBucket is called when exiting the significantBucket production.
	ExitSignificantBucket(c *SignificantBucketContext)

	// ExitFiltersBucket is called when exiting the filtersBucket production.
	ExitFiltersBucket(c *FiltersBucketContext)

	// ExitSparkStatement is called when exiting the sparkStatement production.
	ExitSparkStatement(c *SparkStatementContext)

	// ExitDataStruct is called when exiting the dataStruct production.
	ExitDataStruct(c *DataStructContext)

	// ExitJoinQueryAnalysisStatement is called when exiting the joinQueryAnalysisStatement production.
	ExitJoinQueryAnalysisStatement(c *JoinQueryAnalysisStatementContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitParamValues is called when exiting the paramValues production.
	ExitParamValues(c *ParamValuesContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitUpdateReplaceField is called when exiting the updateReplaceField production.
	ExitUpdateReplaceField(c *UpdateReplaceFieldContext)

	// ExitUpdateAddField is called when exiting the updateAddField production.
	ExitUpdateAddField(c *UpdateAddFieldContext)

	// ExitUpdateRemoveField is called when exiting the updateRemoveField production.
	ExitUpdateRemoveField(c *UpdateRemoveFieldContext)

	// ExitBatchUpdateStatement is called when exiting the batchUpdateStatement production.
	ExitBatchUpdateStatement(c *BatchUpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitDescStatement is called when exiting the descStatement production.
	ExitDescStatement(c *DescStatementContext)

	// ExitAddAlias is called when exiting the addAlias production.
	ExitAddAlias(c *AddAliasContext)

	// ExitDeleteAlias is called when exiting the deleteAlias production.
	ExitDeleteAlias(c *DeleteAliasContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitHive2Statement is called when exiting the hive2Statement production.
	ExitHive2Statement(c *Hive2StatementContext)

	// ExitJdbcStatement is called when exiting the jdbcStatement production.
	ExitJdbcStatement(c *JdbcStatementContext)

	// ExitBasicSQL is called when exiting the basicSQL production.
	ExitBasicSQL(c *BasicSQLContext)

	// ExitFileLoadStatement is called when exiting the fileLoadStatement production.
	ExitFileLoadStatement(c *FileLoadStatementContext)

	// ExitFieldDefine is called when exiting the fieldDefine production.
	ExitFieldDefine(c *FieldDefineContext)

	// ExitAnalyzeStatement is called when exiting the analyzeStatement production.
	ExitAnalyzeStatement(c *AnalyzeStatementContext)

	// ExitHanLPStatement is called when exiting the hanLPStatement production.
	ExitHanLPStatement(c *HanLPStatementContext)

	// ExitNlpFunc is called when exiting the nlpFunc production.
	ExitNlpFunc(c *NlpFuncContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitPathIdentifier is called when exiting the pathIdentifier production.
	ExitPathIdentifier(c *PathIdentifierContext)

	// ExitFieldIdentifier is called when exiting the fieldIdentifier production.
	ExitFieldIdentifier(c *FieldIdentifierContext)

	// ExitIndexIdentifier is called when exiting the indexIdentifier production.
	ExitIndexIdentifier(c *IndexIdentifierContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitStrictIdentifier is called when exiting the strictIdentifier production.
	ExitStrictIdentifier(c *StrictIdentifierContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitNonReserved is called when exiting the nonReserved production.
	ExitNonReserved(c *NonReservedContext)
}
