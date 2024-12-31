package elasticsql

import "testing"

func TestMetricDistinct(t *testing.T){
	sql:="select * from simba_online where ctime>'now-1d' | distinct() map r1"
	RunAndPrintResult(sql, "r1")
}