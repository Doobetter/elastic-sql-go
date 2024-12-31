package agg

import (
	"errors"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/query"
)
//AnalysisStatement 包含多个 AggregationStatement
type AnalysisStatement struct {
	frontStat *query.QueryStatement
	front string // frontStat的名字
	aggStatements []*AggregationStatement
}

func (a *AnalysisStatement) Init(ctx *basic.ExeElasticSQLCtx) error {
	if a.frontStat == nil{
		return errors.New("front statement is not set in analysis statement")
	}
	for _, stat := range a.aggStatements {
		stat.GenPostProcessCode()
		stat.AddAggregationBuilder(a.frontStat.SearchSource)
	}
	return  nil
}

func (a *AnalysisStatement) Execute(ctx *basic.ExeElasticSQLCtx) error {
	// 获取结果，查询统计已经在QueryStatement的执行中实现
	fronResult := ctx.GetResultSet(a.front)
	sr:=fronResult.SearchResponse
	for _, agg := range a.aggStatements {
		ctx.AddResultSet(agg.Name, agg.ParseResult(sr))
	}

	return nil
}

func (a AnalysisStatement) Clean(ctx *basic.ExeElasticSQLCtx) {
	// do nothing
}

func (a AnalysisStatement) GetExportFileName() string {
	return ""
}

func (a *AnalysisStatement) DefaultNameSuffix() string {
	return "_analysis"
}

func (a AnalysisStatement) GetName() string {
	// do nothing
	return ""
}

func (a AnalysisStatement) SetName(name string) {
	// do nothing
}

func (a AnalysisStatement) GenPostProcessCode() int {
	// do nothing
	return -1
}

