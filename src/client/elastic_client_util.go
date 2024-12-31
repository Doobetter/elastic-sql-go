package client

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type ElasticClientUtil struct {
	ctx context.Context
}
func NewElasticClientUtil(){
	util:=new(ElasticClientUtil)
	util.ctx = context.Background()
}
func (e* ElasticClientUtil)OfflineIndex(client *elastic.Client, indexName string)bool{
	response,err:=client.CloseIndex(indexName).Do(e.ctx)
	if err!=nil{
		//
	}
	return response.Acknowledged
}
