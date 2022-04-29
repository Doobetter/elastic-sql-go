package agg

import (
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/olivere/elastic/v7"
	"strings"
)

type AggregationAdapter interface{
	 IsBucket() bool
	 GenAggName() string
}

const (
	NESTED_AGG_NAME_PREFIX = "agg_nested_"
	REVERSE_NESTED_AGG_NAME = "agg_reverse_nested"
)

type BucketAggregationAdapter struct {
	AggName string
	Field string
	Paths string
	PathStr string
	ParentPathStr string
	Order int
	Next AggregationAdapter
	RelatedAgg AggregationAdapter
}

func (b *BucketAggregationAdapter) IsBucket() bool {
	return true
}
func  (b *BucketAggregationAdapter)  GenAggName() string {
	return ""
}

func (b *BucketAggregationAdapter) ToAggregationBuilder() elastic.Aggregation {
	return nil
}

type TermsAggregationAdapter struct {
	BucketAggregationAdapter
	minDocCount int // default is 1
	topN int // default elastic.sql.agg.grouping.default.size
	orderPath string
	include []interface{}
	exclude []interface{}
	asc bool // default desc
	missing interface{}
	script *grammer.ScriptAdapter
}
func (b *TermsAggregationAdapter) ToAggregationBuilder() (elastic.Aggregation,error) {
	aggBuilder := elastic.NewTermsAggregation()
	if b.script != nil{
		sc,err:=b.script.ToScript()
		if err!=nil{
			return nil,err
		}
		aggBuilder.Script(sc)
	}else {
		aggBuilder.Field(b.Field)
	}

	if len(b.include)>0{
		aggBuilder.IncludeValues(b.include...)
	}
	if len(b.exclude)>0{
		aggBuilder.ExcludeValues(b.exclude...)
	}


	aggBuilder.MinDocCount(b.minDocCount)
	aggBuilder.Size(b.topN)

	if b.missing !=nil{
		aggBuilder.Missing(b.missing)
	}

	if b.orderPath != ""{
		if "_count" == b.orderPath{
			aggBuilder.OrderByCount(b.asc)
		}else if "_key" == b.orderPath{
			aggBuilder.OrderByKey(b.asc)
		}else if "_score" == b.orderPath{
			// todo
		}else{
			aggBuilder.Order(b.orderPath,b.asc)
		}

	}


	return aggBuilder,nil
}

func (b * TermsAggregationAdapter) GenAggName() string  {
	b.AggName = "agg_b_terms_" + strings.ReplaceAll(b.Field,".", "_")
	return b.AggName;
}