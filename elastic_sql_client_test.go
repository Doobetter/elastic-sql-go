package elasticsql


import (
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/common"
	"testing"
)

func RunAndPrintTheResultAtSomeCuster(sql string, conf string){
	elasticSQL,err:=ElasticSQL(sql,conf)
	if err!=nil{
		fmt.Println("error ",err)
		return
	}
	elasticSQL.Init()
	elasticSQL.Execute()
	fmt.Println(common.JSONStr(elasticSQL.GetTheResultSet()))
}

var CMS_CONF ="elastic-sql-rest-cms.yml"
func TestMultiCluster(t *testing.T) {
	sql1:="select * from comosproductv6"
	RunAndPrintTheResultAtSomeCuster(sql1,CMS_CONF)

	sql2:="select * from simba*"
	RunAndPrintTheResultAtSomeCuster(sql2,"")

}


var G1_CONF ="elastic-sql-rest-g1.yml"
func TestV8_2_3(t *testing.T) {
	sql1:="select * from test1"
	RunAndPrintTheResultAtSomeCuster(sql1,G1_CONF)
}