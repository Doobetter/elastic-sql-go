// Code generated from ./src/parser/ElasticSQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ElasticSQL
import "github.com/antlr4-go/antlr/v4"

// BaseElasticSQLListener is a complete listener for a parse tree produced by ElasticSQLParser.
type BaseElasticSQLListener struct{}

var _ ElasticSQLListener = &BaseElasticSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseElasticSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseElasticSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseElasticSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseElasticSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterElasticSQL is called when production elasticSQL is entered.
func (s *BaseElasticSQLListener) EnterElasticSQL(ctx *ElasticSQLContext) {}

// ExitElasticSQL is called when production elasticSQL is exited.
func (s *BaseElasticSQLListener) ExitElasticSQL(ctx *ElasticSQLContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseElasticSQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseElasticSQLListener) ExitStatement(ctx *StatementContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseElasticSQLListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseElasticSQLListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *BaseElasticSQLListener) EnterSelectItem(ctx *SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *BaseElasticSQLListener) ExitSelectItem(ctx *SelectItemContext) {}

// EnterCollapseExpr is called when production collapseExpr is entered.
func (s *BaseElasticSQLListener) EnterCollapseExpr(ctx *CollapseExprContext) {}

// ExitCollapseExpr is called when production collapseExpr is exited.
func (s *BaseElasticSQLListener) ExitCollapseExpr(ctx *CollapseExprContext) {}

// EnterCustomScoreExpr is called when production customScoreExpr is entered.
func (s *BaseElasticSQLListener) EnterCustomScoreExpr(ctx *CustomScoreExprContext) {}

// ExitCustomScoreExpr is called when production customScoreExpr is exited.
func (s *BaseElasticSQLListener) ExitCustomScoreExpr(ctx *CustomScoreExprContext) {}

// EnterRescoreExpr is called when production rescoreExpr is entered.
func (s *BaseElasticSQLListener) EnterRescoreExpr(ctx *RescoreExprContext) {}

// ExitRescoreExpr is called when production rescoreExpr is exited.
func (s *BaseElasticSQLListener) ExitRescoreExpr(ctx *RescoreExprContext) {}

// EnterMemSort is called when production memSort is entered.
func (s *BaseElasticSQLListener) EnterMemSort(ctx *MemSortContext) {}

// ExitMemSort is called when production memSort is exited.
func (s *BaseElasticSQLListener) ExitMemSort(ctx *MemSortContext) {}

// EnterInnerHit is called when production innerHit is entered.
func (s *BaseElasticSQLListener) EnterInnerHit(ctx *InnerHitContext) {}

// ExitInnerHit is called when production innerHit is exited.
func (s *BaseElasticSQLListener) ExitInnerHit(ctx *InnerHitContext) {}

// EnterScriptField is called when production scriptField is entered.
func (s *BaseElasticSQLListener) EnterScriptField(ctx *ScriptFieldContext) {}

// ExitScriptField is called when production scriptField is exited.
func (s *BaseElasticSQLListener) ExitScriptField(ctx *ScriptFieldContext) {}

// EnterHighlight is called when production highlight is entered.
func (s *BaseElasticSQLListener) EnterHighlight(ctx *HighlightContext) {}

// ExitHighlight is called when production highlight is exited.
func (s *BaseElasticSQLListener) ExitHighlight(ctx *HighlightContext) {}

// EnterFieldAs is called when production fieldAs is entered.
func (s *BaseElasticSQLListener) EnterFieldAs(ctx *FieldAsContext) {}

// ExitFieldAs is called when production fieldAs is exited.
func (s *BaseElasticSQLListener) ExitFieldAs(ctx *FieldAsContext) {}

// EnterWhereExpression is called when production whereExpression is entered.
func (s *BaseElasticSQLListener) EnterWhereExpression(ctx *WhereExpressionContext) {}

// ExitWhereExpression is called when production whereExpression is exited.
func (s *BaseElasticSQLListener) ExitWhereExpression(ctx *WhereExpressionContext) {}

// EnterLogicalExpr is called when production logicalExpr is entered.
func (s *BaseElasticSQLListener) EnterLogicalExpr(ctx *LogicalExprContext) {}

// ExitLogicalExpr is called when production logicalExpr is exited.
func (s *BaseElasticSQLListener) ExitLogicalExpr(ctx *LogicalExprContext) {}

// EnterComparableExpression is called when production comparableExpression is entered.
func (s *BaseElasticSQLListener) EnterComparableExpression(ctx *ComparableExpressionContext) {}

// ExitComparableExpression is called when production comparableExpression is exited.
func (s *BaseElasticSQLListener) ExitComparableExpression(ctx *ComparableExpressionContext) {}

// EnterTermCompare is called when production termCompare is entered.
func (s *BaseElasticSQLListener) EnterTermCompare(ctx *TermCompareContext) {}

// ExitTermCompare is called when production termCompare is exited.
func (s *BaseElasticSQLListener) ExitTermCompare(ctx *TermCompareContext) {}

// EnterBtwCompare is called when production btwCompare is entered.
func (s *BaseElasticSQLListener) EnterBtwCompare(ctx *BtwCompareContext) {}

// ExitBtwCompare is called when production btwCompare is exited.
func (s *BaseElasticSQLListener) ExitBtwCompare(ctx *BtwCompareContext) {}

// EnterArithmeticExpressionCompare is called when production arithmeticExpressionCompare is entered.
func (s *BaseElasticSQLListener) EnterArithmeticExpressionCompare(ctx *ArithmeticExpressionCompareContext) {
}

// ExitArithmeticExpressionCompare is called when production arithmeticExpressionCompare is exited.
func (s *BaseElasticSQLListener) ExitArithmeticExpressionCompare(ctx *ArithmeticExpressionCompareContext) {
}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseElasticSQLListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseElasticSQLListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterAddition is called when production addition is entered.
func (s *BaseElasticSQLListener) EnterAddition(ctx *AdditionContext) {}

// ExitAddition is called when production addition is exited.
func (s *BaseElasticSQLListener) ExitAddition(ctx *AdditionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BaseElasticSQLListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BaseElasticSQLListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterMulti is called when production multi is entered.
func (s *BaseElasticSQLListener) EnterMulti(ctx *MultiContext) {}

// ExitMulti is called when production multi is exited.
func (s *BaseElasticSQLListener) ExitMulti(ctx *MultiContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseElasticSQLListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseElasticSQLListener) ExitAtom(ctx *AtomContext) {}

// EnterFunctionalCompare is called when production functionalCompare is entered.
func (s *BaseElasticSQLListener) EnterFunctionalCompare(ctx *FunctionalCompareContext) {}

// ExitFunctionalCompare is called when production functionalCompare is exited.
func (s *BaseElasticSQLListener) ExitFunctionalCompare(ctx *FunctionalCompareContext) {}

// EnterJoinFunction is called when production joinFunction is entered.
func (s *BaseElasticSQLListener) EnterJoinFunction(ctx *JoinFunctionContext) {}

// ExitJoinFunction is called when production joinFunction is exited.
func (s *BaseElasticSQLListener) ExitJoinFunction(ctx *JoinFunctionContext) {}

// EnterJoinFunctionNames is called when production joinFunctionNames is entered.
func (s *BaseElasticSQLListener) EnterJoinFunctionNames(ctx *JoinFunctionNamesContext) {}

// ExitJoinFunctionNames is called when production joinFunctionNames is exited.
func (s *BaseElasticSQLListener) ExitJoinFunctionNames(ctx *JoinFunctionNamesContext) {}

// EnterScriptFunction is called when production scriptFunction is entered.
func (s *BaseElasticSQLListener) EnterScriptFunction(ctx *ScriptFunctionContext) {}

// ExitScriptFunction is called when production scriptFunction is exited.
func (s *BaseElasticSQLListener) ExitScriptFunction(ctx *ScriptFunctionContext) {}

// EnterFullLevelFunction is called when production fullLevelFunction is entered.
func (s *BaseElasticSQLListener) EnterFullLevelFunction(ctx *FullLevelFunctionContext) {}

// ExitFullLevelFunction is called when production fullLevelFunction is exited.
func (s *BaseElasticSQLListener) ExitFullLevelFunction(ctx *FullLevelFunctionContext) {}

// EnterFullLevelFunctionNames is called when production fullLevelFunctionNames is entered.
func (s *BaseElasticSQLListener) EnterFullLevelFunctionNames(ctx *FullLevelFunctionNamesContext) {}

// ExitFullLevelFunctionNames is called when production fullLevelFunctionNames is exited.
func (s *BaseElasticSQLListener) ExitFullLevelFunctionNames(ctx *FullLevelFunctionNamesContext) {}

// EnterTermLevelFunction is called when production termLevelFunction is entered.
func (s *BaseElasticSQLListener) EnterTermLevelFunction(ctx *TermLevelFunctionContext) {}

// ExitTermLevelFunction is called when production termLevelFunction is exited.
func (s *BaseElasticSQLListener) ExitTermLevelFunction(ctx *TermLevelFunctionContext) {}

// EnterTermLevelFunctionNames is called when production termLevelFunctionNames is entered.
func (s *BaseElasticSQLListener) EnterTermLevelFunctionNames(ctx *TermLevelFunctionNamesContext) {}

// ExitTermLevelFunctionNames is called when production termLevelFunctionNames is exited.
func (s *BaseElasticSQLListener) ExitTermLevelFunctionNames(ctx *TermLevelFunctionNamesContext) {}

// EnterProp is called when production prop is entered.
func (s *BaseElasticSQLListener) EnterProp(ctx *PropContext) {}

// ExitProp is called when production prop is exited.
func (s *BaseElasticSQLListener) ExitProp(ctx *PropContext) {}

// EnterParam is called when production param is entered.
func (s *BaseElasticSQLListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseElasticSQLListener) ExitParam(ctx *ParamContext) {}

// EnterParam2 is called when production param2 is entered.
func (s *BaseElasticSQLListener) EnterParam2(ctx *Param2Context) {}

// ExitParam2 is called when production param2 is exited.
func (s *BaseElasticSQLListener) ExitParam2(ctx *Param2Context) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseElasticSQLListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseElasticSQLListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseElasticSQLListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseElasticSQLListener) ExitSortItem(ctx *SortItemContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseElasticSQLListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseElasticSQLListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterExportField is called when production exportField is entered.
func (s *BaseElasticSQLListener) EnterExportField(ctx *ExportFieldContext) {}

// ExitExportField is called when production exportField is exited.
func (s *BaseElasticSQLListener) ExitExportField(ctx *ExportFieldContext) {}

// EnterAnalysisStatement is called when production analysisStatement is entered.
func (s *BaseElasticSQLListener) EnterAnalysisStatement(ctx *AnalysisStatementContext) {}

// ExitAnalysisStatement is called when production analysisStatement is exited.
func (s *BaseElasticSQLListener) ExitAnalysisStatement(ctx *AnalysisStatementContext) {}

// EnterAggStatement is called when production aggStatement is entered.
func (s *BaseElasticSQLListener) EnterAggStatement(ctx *AggStatementContext) {}

// ExitAggStatement is called when production aggStatement is exited.
func (s *BaseElasticSQLListener) ExitAggStatement(ctx *AggStatementContext) {}

// EnterMetricAgg is called when production metricAgg is entered.
func (s *BaseElasticSQLListener) EnterMetricAgg(ctx *MetricAggContext) {}

// ExitMetricAgg is called when production metricAgg is exited.
func (s *BaseElasticSQLListener) ExitMetricAgg(ctx *MetricAggContext) {}

// EnterMetricNames is called when production metricNames is entered.
func (s *BaseElasticSQLListener) EnterMetricNames(ctx *MetricNamesContext) {}

// ExitMetricNames is called when production metricNames is exited.
func (s *BaseElasticSQLListener) ExitMetricNames(ctx *MetricNamesContext) {}

// EnterMetricParams is called when production metricParams is entered.
func (s *BaseElasticSQLListener) EnterMetricParams(ctx *MetricParamsContext) {}

// ExitMetricParams is called when production metricParams is exited.
func (s *BaseElasticSQLListener) ExitMetricParams(ctx *MetricParamsContext) {}

// EnterMkv is called when production mkv is entered.
func (s *BaseElasticSQLListener) EnterMkv(ctx *MkvContext) {}

// ExitMkv is called when production mkv is exited.
func (s *BaseElasticSQLListener) ExitMkv(ctx *MkvContext) {}

// EnterMetricParamNames is called when production metricParamNames is entered.
func (s *BaseElasticSQLListener) EnterMetricParamNames(ctx *MetricParamNamesContext) {}

// ExitMetricParamNames is called when production metricParamNames is exited.
func (s *BaseElasticSQLListener) ExitMetricParamNames(ctx *MetricParamNamesContext) {}

// EnterScriptPhrase is called when production scriptPhrase is entered.
func (s *BaseElasticSQLListener) EnterScriptPhrase(ctx *ScriptPhraseContext) {}

// ExitScriptPhrase is called when production scriptPhrase is exited.
func (s *BaseElasticSQLListener) ExitScriptPhrase(ctx *ScriptPhraseContext) {}

// EnterBucketAggList is called when production bucketAggList is entered.
func (s *BaseElasticSQLListener) EnterBucketAggList(ctx *BucketAggListContext) {}

// ExitBucketAggList is called when production bucketAggList is exited.
func (s *BaseElasticSQLListener) ExitBucketAggList(ctx *BucketAggListContext) {}

// EnterBucketAgg is called when production bucketAgg is entered.
func (s *BaseElasticSQLListener) EnterBucketAgg(ctx *BucketAggContext) {}

// ExitBucketAgg is called when production bucketAgg is exited.
func (s *BaseElasticSQLListener) ExitBucketAgg(ctx *BucketAggContext) {}

// EnterBucketAggChoice is called when production bucketAggChoice is entered.
func (s *BaseElasticSQLListener) EnterBucketAggChoice(ctx *BucketAggChoiceContext) {}

// ExitBucketAggChoice is called when production bucketAggChoice is exited.
func (s *BaseElasticSQLListener) ExitBucketAggChoice(ctx *BucketAggChoiceContext) {}

// EnterTermsBucket is called when production termsBucket is entered.
func (s *BaseElasticSQLListener) EnterTermsBucket(ctx *TermsBucketContext) {}

// ExitTermsBucket is called when production termsBucket is exited.
func (s *BaseElasticSQLListener) ExitTermsBucket(ctx *TermsBucketContext) {}

// EnterHavingExpr is called when production havingExpr is entered.
func (s *BaseElasticSQLListener) EnterHavingExpr(ctx *HavingExprContext) {}

// ExitHavingExpr is called when production havingExpr is exited.
func (s *BaseElasticSQLListener) ExitHavingExpr(ctx *HavingExprContext) {}

// EnterRangeBucket is called when production rangeBucket is entered.
func (s *BaseElasticSQLListener) EnterRangeBucket(ctx *RangeBucketContext) {}

// ExitRangeBucket is called when production rangeBucket is exited.
func (s *BaseElasticSQLListener) ExitRangeBucket(ctx *RangeBucketContext) {}

// EnterRangeExpr is called when production rangeExpr is entered.
func (s *BaseElasticSQLListener) EnterRangeExpr(ctx *RangeExprContext) {}

// ExitRangeExpr is called when production rangeExpr is exited.
func (s *BaseElasticSQLListener) ExitRangeExpr(ctx *RangeExprContext) {}

// EnterRangeUnit is called when production rangeUnit is entered.
func (s *BaseElasticSQLListener) EnterRangeUnit(ctx *RangeUnitContext) {}

// ExitRangeUnit is called when production rangeUnit is exited.
func (s *BaseElasticSQLListener) ExitRangeUnit(ctx *RangeUnitContext) {}

// EnterRangeFromTo is called when production rangeFromTo is entered.
func (s *BaseElasticSQLListener) EnterRangeFromTo(ctx *RangeFromToContext) {}

// ExitRangeFromTo is called when production rangeFromTo is exited.
func (s *BaseElasticSQLListener) ExitRangeFromTo(ctx *RangeFromToContext) {}

// EnterDateRangeBucket is called when production dateRangeBucket is entered.
func (s *BaseElasticSQLListener) EnterDateRangeBucket(ctx *DateRangeBucketContext) {}

// ExitDateRangeBucket is called when production dateRangeBucket is exited.
func (s *BaseElasticSQLListener) ExitDateRangeBucket(ctx *DateRangeBucketContext) {}

// EnterDateRangeExpr is called when production dateRangeExpr is entered.
func (s *BaseElasticSQLListener) EnterDateRangeExpr(ctx *DateRangeExprContext) {}

// ExitDateRangeExpr is called when production dateRangeExpr is exited.
func (s *BaseElasticSQLListener) ExitDateRangeExpr(ctx *DateRangeExprContext) {}

// EnterDateRange is called when production dateRange is entered.
func (s *BaseElasticSQLListener) EnterDateRange(ctx *DateRangeContext) {}

// ExitDateRange is called when production dateRange is exited.
func (s *BaseElasticSQLListener) ExitDateRange(ctx *DateRangeContext) {}

// EnterDateRangeFromTo is called when production dateRangeFromTo is entered.
func (s *BaseElasticSQLListener) EnterDateRangeFromTo(ctx *DateRangeFromToContext) {}

// ExitDateRangeFromTo is called when production dateRangeFromTo is exited.
func (s *BaseElasticSQLListener) ExitDateRangeFromTo(ctx *DateRangeFromToContext) {}

// EnterHistogramBucket is called when production histogramBucket is entered.
func (s *BaseElasticSQLListener) EnterHistogramBucket(ctx *HistogramBucketContext) {}

// ExitHistogramBucket is called when production histogramBucket is exited.
func (s *BaseElasticSQLListener) ExitHistogramBucket(ctx *HistogramBucketContext) {}

// EnterDateHistogramBucket is called when production dateHistogramBucket is entered.
func (s *BaseElasticSQLListener) EnterDateHistogramBucket(ctx *DateHistogramBucketContext) {}

// ExitDateHistogramBucket is called when production dateHistogramBucket is exited.
func (s *BaseElasticSQLListener) ExitDateHistogramBucket(ctx *DateHistogramBucketContext) {}

// EnterSignificantBucket is called when production significantBucket is entered.
func (s *BaseElasticSQLListener) EnterSignificantBucket(ctx *SignificantBucketContext) {}

// ExitSignificantBucket is called when production significantBucket is exited.
func (s *BaseElasticSQLListener) ExitSignificantBucket(ctx *SignificantBucketContext) {}

// EnterFiltersBucket is called when production filtersBucket is entered.
func (s *BaseElasticSQLListener) EnterFiltersBucket(ctx *FiltersBucketContext) {}

// ExitFiltersBucket is called when production filtersBucket is exited.
func (s *BaseElasticSQLListener) ExitFiltersBucket(ctx *FiltersBucketContext) {}

// EnterSparkStatement is called when production sparkStatement is entered.
func (s *BaseElasticSQLListener) EnterSparkStatement(ctx *SparkStatementContext) {}

// ExitSparkStatement is called when production sparkStatement is exited.
func (s *BaseElasticSQLListener) ExitSparkStatement(ctx *SparkStatementContext) {}

// EnterDataStruct is called when production dataStruct is entered.
func (s *BaseElasticSQLListener) EnterDataStruct(ctx *DataStructContext) {}

// ExitDataStruct is called when production dataStruct is exited.
func (s *BaseElasticSQLListener) ExitDataStruct(ctx *DataStructContext) {}

// EnterJoinQueryAnalysisStatement is called when production joinQueryAnalysisStatement is entered.
func (s *BaseElasticSQLListener) EnterJoinQueryAnalysisStatement(ctx *JoinQueryAnalysisStatementContext) {
}

// ExitJoinQueryAnalysisStatement is called when production joinQueryAnalysisStatement is exited.
func (s *BaseElasticSQLListener) ExitJoinQueryAnalysisStatement(ctx *JoinQueryAnalysisStatementContext) {
}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseElasticSQLListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseElasticSQLListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseElasticSQLListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseElasticSQLListener) ExitFieldList(ctx *FieldListContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseElasticSQLListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseElasticSQLListener) ExitValueList(ctx *ValueListContext) {}

// EnterParamValues is called when production paramValues is entered.
func (s *BaseElasticSQLListener) EnterParamValues(ctx *ParamValuesContext) {}

// ExitParamValues is called when production paramValues is exited.
func (s *BaseElasticSQLListener) ExitParamValues(ctx *ParamValuesContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseElasticSQLListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseElasticSQLListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterUpdateReplaceField is called when production updateReplaceField is entered.
func (s *BaseElasticSQLListener) EnterUpdateReplaceField(ctx *UpdateReplaceFieldContext) {}

// ExitUpdateReplaceField is called when production updateReplaceField is exited.
func (s *BaseElasticSQLListener) ExitUpdateReplaceField(ctx *UpdateReplaceFieldContext) {}

// EnterUpdateAddField is called when production updateAddField is entered.
func (s *BaseElasticSQLListener) EnterUpdateAddField(ctx *UpdateAddFieldContext) {}

// ExitUpdateAddField is called when production updateAddField is exited.
func (s *BaseElasticSQLListener) ExitUpdateAddField(ctx *UpdateAddFieldContext) {}

// EnterUpdateRemoveField is called when production updateRemoveField is entered.
func (s *BaseElasticSQLListener) EnterUpdateRemoveField(ctx *UpdateRemoveFieldContext) {}

// ExitUpdateRemoveField is called when production updateRemoveField is exited.
func (s *BaseElasticSQLListener) ExitUpdateRemoveField(ctx *UpdateRemoveFieldContext) {}

// EnterBatchUpdateStatement is called when production batchUpdateStatement is entered.
func (s *BaseElasticSQLListener) EnterBatchUpdateStatement(ctx *BatchUpdateStatementContext) {}

// ExitBatchUpdateStatement is called when production batchUpdateStatement is exited.
func (s *BaseElasticSQLListener) ExitBatchUpdateStatement(ctx *BatchUpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseElasticSQLListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseElasticSQLListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterDescStatement is called when production descStatement is entered.
func (s *BaseElasticSQLListener) EnterDescStatement(ctx *DescStatementContext) {}

// ExitDescStatement is called when production descStatement is exited.
func (s *BaseElasticSQLListener) ExitDescStatement(ctx *DescStatementContext) {}

// EnterAddAlias is called when production addAlias is entered.
func (s *BaseElasticSQLListener) EnterAddAlias(ctx *AddAliasContext) {}

// ExitAddAlias is called when production addAlias is exited.
func (s *BaseElasticSQLListener) ExitAddAlias(ctx *AddAliasContext) {}

// EnterDeleteAlias is called when production deleteAlias is entered.
func (s *BaseElasticSQLListener) EnterDeleteAlias(ctx *DeleteAliasContext) {}

// ExitDeleteAlias is called when production deleteAlias is exited.
func (s *BaseElasticSQLListener) ExitDeleteAlias(ctx *DeleteAliasContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseElasticSQLListener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseElasticSQLListener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterHive2Statement is called when production hive2Statement is entered.
func (s *BaseElasticSQLListener) EnterHive2Statement(ctx *Hive2StatementContext) {}

// ExitHive2Statement is called when production hive2Statement is exited.
func (s *BaseElasticSQLListener) ExitHive2Statement(ctx *Hive2StatementContext) {}

// EnterJdbcStatement is called when production jdbcStatement is entered.
func (s *BaseElasticSQLListener) EnterJdbcStatement(ctx *JdbcStatementContext) {}

// ExitJdbcStatement is called when production jdbcStatement is exited.
func (s *BaseElasticSQLListener) ExitJdbcStatement(ctx *JdbcStatementContext) {}

// EnterBasicSQL is called when production basicSQL is entered.
func (s *BaseElasticSQLListener) EnterBasicSQL(ctx *BasicSQLContext) {}

// ExitBasicSQL is called when production basicSQL is exited.
func (s *BaseElasticSQLListener) ExitBasicSQL(ctx *BasicSQLContext) {}

// EnterFileLoadStatement is called when production fileLoadStatement is entered.
func (s *BaseElasticSQLListener) EnterFileLoadStatement(ctx *FileLoadStatementContext) {}

// ExitFileLoadStatement is called when production fileLoadStatement is exited.
func (s *BaseElasticSQLListener) ExitFileLoadStatement(ctx *FileLoadStatementContext) {}

// EnterFieldDefine is called when production fieldDefine is entered.
func (s *BaseElasticSQLListener) EnterFieldDefine(ctx *FieldDefineContext) {}

// ExitFieldDefine is called when production fieldDefine is exited.
func (s *BaseElasticSQLListener) ExitFieldDefine(ctx *FieldDefineContext) {}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseElasticSQLListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseElasticSQLListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterHanLPStatement is called when production hanLPStatement is entered.
func (s *BaseElasticSQLListener) EnterHanLPStatement(ctx *HanLPStatementContext) {}

// ExitHanLPStatement is called when production hanLPStatement is exited.
func (s *BaseElasticSQLListener) ExitHanLPStatement(ctx *HanLPStatementContext) {}

// EnterNlpFunc is called when production nlpFunc is entered.
func (s *BaseElasticSQLListener) EnterNlpFunc(ctx *NlpFuncContext) {}

// ExitNlpFunc is called when production nlpFunc is exited.
func (s *BaseElasticSQLListener) ExitNlpFunc(ctx *NlpFuncContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseElasticSQLListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseElasticSQLListener) ExitDataType(ctx *DataTypeContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseElasticSQLListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseElasticSQLListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterPathIdentifier is called when production pathIdentifier is entered.
func (s *BaseElasticSQLListener) EnterPathIdentifier(ctx *PathIdentifierContext) {}

// ExitPathIdentifier is called when production pathIdentifier is exited.
func (s *BaseElasticSQLListener) ExitPathIdentifier(ctx *PathIdentifierContext) {}

// EnterFieldIdentifier is called when production fieldIdentifier is entered.
func (s *BaseElasticSQLListener) EnterFieldIdentifier(ctx *FieldIdentifierContext) {}

// ExitFieldIdentifier is called when production fieldIdentifier is exited.
func (s *BaseElasticSQLListener) ExitFieldIdentifier(ctx *FieldIdentifierContext) {}

// EnterIndexIdentifier is called when production indexIdentifier is entered.
func (s *BaseElasticSQLListener) EnterIndexIdentifier(ctx *IndexIdentifierContext) {}

// ExitIndexIdentifier is called when production indexIdentifier is exited.
func (s *BaseElasticSQLListener) ExitIndexIdentifier(ctx *IndexIdentifierContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseElasticSQLListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseElasticSQLListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterStrictIdentifier is called when production strictIdentifier is entered.
func (s *BaseElasticSQLListener) EnterStrictIdentifier(ctx *StrictIdentifierContext) {}

// ExitStrictIdentifier is called when production strictIdentifier is exited.
func (s *BaseElasticSQLListener) ExitStrictIdentifier(ctx *StrictIdentifierContext) {}

// EnterStr is called when production str is entered.
func (s *BaseElasticSQLListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseElasticSQLListener) ExitStr(ctx *StrContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseElasticSQLListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseElasticSQLListener) ExitNumber(ctx *NumberContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseElasticSQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseElasticSQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseElasticSQLListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseElasticSQLListener) ExitNonReserved(ctx *NonReservedContext) {}
