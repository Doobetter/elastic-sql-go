package agg

import (
	"github.com/olivere/elastic/v7"
	"strings"
)

type MetricAggregationAdapterI interface {
	// GetMetricValue 获取结果值
	GetMetricValue(aggOfParent elastic.Aggregations) interface{}
	// GetSchemaName 获取输出字段名
	GetSchemaName()string
	Next() MetricAggregationAdapterI
}

type MetricAggregationAdapter struct {
	AggName string
	Field string
	metricName string
	metricAS string
	metricParam map[string]interface{}
	next MetricAggregationAdapterI
}

func (a * MetricAggregationAdapter) Next()MetricAggregationAdapterI{
	return a.next
}

func (a * MetricAggregationAdapter) GetSchemaName()string{
	if a.metricAS == ""{
		return a.metricName
	}
	return a.metricAS
}

func (a * MetricAggregationAdapter) GenAggName() string  {

	if a.metricAS != ""{
		a.AggName = a.metricName
	}else {
		a.AggName = "agg_m_" + a.metricName + "_" + strings.ReplaceAll(a.Field,".","_")
	}
	return a.AggName;
}
func (a *MetricAggregationAdapter) IsBucket() bool {
	return false
}
func (a *MetricAggregationAdapter) GetMetricValue(aggOfParent elastic.Aggregations) interface{} {
	return nil
}



// ParseMetricResult 解析
func  ParseMetricResult(row map[string]interface{},adapter MetricAggregationAdapterI, aggOfParent elastic.Aggregations)  {
	metric:=adapter
	for true {
		if metric == nil {
			break
		}
		row[metric.GetSchemaName()] = metric.GetMetricValue(aggOfParent)
		metric = metric.Next()
	}
}



func GetInstanceFromMetricName(metricName string) MetricAggregationAdapterI{
	switch metricName {
	case "sum" :
			return new(SumAggregationAdapter)
	}
	return nil
}


type SumAggregationAdapter struct {
	MetricAggregationAdapter
}
func (a *SumAggregationAdapter) ToAggregationBuilder() elastic.Aggregation {
	return elastic.NewSumAggregation().Field(a.Field)

}

func (a * SumAggregationAdapter) GetMetricValue(aggOfParent elastic.Aggregations) interface{} {
	sum,found:=aggOfParent.Sum(a.AggName)
	if found == false{
		return float32(0)
	}else{
		return sum.Value
	}
}
