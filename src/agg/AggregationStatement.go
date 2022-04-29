package agg

import (
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/olivere/elastic/v7"
)

type AggregationStatement struct {
	basic.ProcessUnit
	Where    grammer.Expression // filter aggregation

	firstAgg AggregationAdapter
	Metrics []*MetricAggregationAdapter
	Buckets []*BucketAggregationAdapter
	topN int  // 获取桶数
	Export  *grammer.ExportClause

}


// AddAggregationBuilder 将adapter写入SearchSourceBuilder
func (a *AggregationStatement)AddAggregationBuilder(searchSource *elastic.SearchSource){
	//var lastBucket elastic.Aggregations
	if len(a.Buckets)>0{
		 bucket := a.Buckets[0]
		 searchSource.Aggregation( bucket.AggName, bucket.ToAggregationBuilder())
	}
}

func (a* AggregationStatement) ParseResult(searchResult *elastic.SearchResult) *basic.ResultSet {
	resultSet := basic.NewResultSet(a.Name)

	aggs:=  searchResult.Aggregations

	total := searchResult.TotalHits()
	if total <=0 || aggs == nil{
		resultSet.ErrMsg = "QUERY记录数为0, 没有AGGS"
		resultSet.Count = 0
		return resultSet
	}

	var rows []map[string]interface{}


	if a.firstAgg.IsBucket(){
		// TODO 判断是否是filteAgg

	}else {
		row:= make(map[string]interface{})
		adapter:=a.firstAgg.(*MetricAggregationAdapter)
		ParseMetricResult(row,adapter,aggs)
		if len(row) >0{
			rows = append(rows, row)
		}
	}

	return resultSet

}