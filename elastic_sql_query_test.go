package elasticsql

import (
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/common"
	"testing"
)

func TestDriver(t *testing.T) {
	sql :="select * from wb_tblog_all_20200901"
	RunAndPrintTheResult(sql)
}

func RunAndPrintTheResult(sql string){
	elasticSQL:=ElasticSQL(sql,"")
	elasticSQL.Init()
	elasticSQL.Execute()
	fmt.Println(common.JSONPrettyStr(elasticSQL.GetTheResultSet()))
}
func RunAndPrintTheResultInSomeCluster(sql string,conf string){
	elasticSQL:=ElasticSQL(sql,conf)
	elasticSQL.Init()
	elasticSQL.Execute()
	fmt.Println(common.JSONPrettyStr(elasticSQL.GetTheResultSet()))
}

func TestFields(t *testing.T) {
	//sql :="select _id from wb_tblog_all_20200901"
	//sql :="select _id,_index from wb_tblog_all_20200901"
	sql :="select _id,nick,ctime from wb_tblog_all_20200901"
	RunAndPrintTheResult(sql)
}

func TestLimit(t *testing.T) {
	sql :="select _id,nick,ctime from wb_tblog_all_20200901 limit  20 "
	RunAndPrintTheResult(sql)
}

func TestOrder(t *testing.T) {
	sql :="select _id,ctime from simba_online order by ctime desc "
	RunAndPrintTheResult(sql)
}

func TestMap(t *testing.T) {
	sql :="select _id,ctime from simba_online limit  20  map t1"
	RunAndPrintTheResult(sql)
}

func TestTermExpression(t *testing.T){
	sql := "select _id,docType from simba_online where docType='news'"
	RunAndPrintTheResult(sql)
}
func TestTermExpression2(t *testing.T){
	sql := "select _id,docType from simba_online where ctime>1632451800"
	RunAndPrintTheResult(sql)
}
func TestTermExpressionAnd(t *testing.T){
	sql := "select _id,docType from simba_online where ctime>1632451800 and ctime>1632451802 and ctime>1632451801 "
	RunAndPrintTheResult(sql)
}
func TestTermExpressionOR(t *testing.T){
	sql := "select _id,docType from simba_online where ctime>1632451800 or ctime>1632451802 OR ctime>1632451801 "
	RunAndPrintTheResult(sql)
}
func TestTermExpressionAndOR(t *testing.T){
	sql := "select _id,docType from simba_online where ctime> 1632451800  or ctime>1632451800 or (ctime>1632451802 and ctime>1632451801 and ctime>1632451803) "
	RunAndPrintTheResult(sql)
}
func TestAndOR(t *testing.T){
	sql := "select _id,docType from simba_online where appEtime>1631531930 and area_tqt$text_area$province='北京' and toutiaoIdx=1 and area_tqt$distribute_features$is_regional=True"
	RunAndPrintTheResult(sql)
}

func TestAndOR2(t *testing.T){
	sql := "select _id,docType from simba_online where appEtime>1631531930 and area_tqt$text_area$province='北京'  and area_tqt$distribute_features$is_regional=True or (toutiaoIdx!=1 or docType='subject') "
	RunAndPrintTheResult(sql)
}

func TestBtw(t *testing.T){
	sql := "select _id,docType from simba_online where appEtime between 1631531930 and 1632451801"
	RunAndPrintTheResult(sql)
}
func TestBtw2(t *testing.T){
	sql := "select _id,docType from simba_online where appEtime range[ 1631531930 , 1632451801) and (docType!='news')"
	RunAndPrintTheResult(sql)
}

func TestFunctionInOutMissExist(t *testing.T){
	sql := "select _id,docType from simba_online where not docType in('video','news')"
	RunAndPrintTheResult(sql)
}

func TestFunctionHasAll_HasAny(t *testing.T){
	//sql := "select _id,sni_type from simba_online where sni_type has_all(55,105)"
	sql := "select _id,sni_type from simba_online where sni_type has_any(55,105)"
	RunAndPrintTheResult(sql)
}

func TestFunctionLike_Rlike(t *testing.T){
	sql := "select _id,labels from simba_online where labels rlike('股.*')"
	//sql := "select _id,labels from simba_online where labels like('股*')"
	RunAndPrintTheResult(sql)
}


func TestFullTextMatch(t *testing.T){
	//sql := "select _id,title from simba_online where  match(field='title',query='字节')"
	sql := "select _id,title,_score from simba_online where ctime>1631609609 and match(field='title',query='字节跳动回应',minimum_should_match='50%',operator='OR',boost=2)"
	RunAndPrintTheResult(sql)
}

func TestFullTextMatchPhrase(t *testing.T){
	sql := "select _id,title from simba_online where  match_phrase(field='title',query='字节跳动回应',slop=1)"
	RunAndPrintTheResult(sql)
}

func TestFullTextQueryString(t *testing.T){
	sql := "select _id,title from simba_online where  not query_string(fields='title',query='字节跳动回应',slop=1)"
	RunAndPrintTheResult(sql)
}
func TestFullTextHighlight(t *testing.T){
	sql := "select _id,highlight(title) by 'b' from simba_online where query_string(fields='title',query='字节跳动回应',slop=1)"
	RunAndPrintTheResult(sql)
}

func TestObject(t *testing.T){
	sql := "select _id,title,subjects.docID from acomos2021 where subjects$docID='kftpnny2234728'"
	RunAndPrintTheResultInSomeCluster(sql,"elastic-sql-rest-cms.yml")
}

func TestNested1(t *testing.T){
	sql := "select _id from acomos2021 where a.b=1"
	RunAndPrintTheResultInSomeCluster(sql,"elastic-sql-rest-cms.yml")
}

func TestNested2(t *testing.T){
	//sql := "select _id from acomos2021 where a.b=1 and a.c > 1 "
	//sql := "select _id from acomos2021 where a.b=1 and a.c > 2 and a.d > 1  "
	//sql := "select _id from acomos2021 where a.b=1 or a.c > 2 or a.d > 1  "
	// sql := "select _id from acomos2021 where a.b=1 or a.c > 2 or a.c > 1  "
	// sql := "select _id from acomos2021 where a.b=1 or a.c > 1 or a.c > 2  "
	// sql := "select _id from acomos2021 where a.b=1 or a.c > 1 or c > 2  "
	//sql := "select _id from acomos2021 where a.b=1 or a.c > 2 and  a.c > 1  "
	sql := "select _id from acomos2021 where a.b=1 or a.c > 2 and  a.d > 1  "
	//sql := "select _id from acomos2021 where b=1 or a.c > 2 and  a.d > 1 "
	RunAndPrintTheResultInSomeCluster(sql,"elastic-sql-rest-cms.yml")
}

func TestNested3(t *testing.T){

	//sql := "select _id from acomos2021 where a.b=1 or (a.c > 2 and  a.d > 1)  "
	sql := "select _id from simba_online where a.c.d.a=1 or a.c.d.b > 2"
	RunAndPrintTheResult(sql)
}

func TestExportJSONByScroll(t *testing.T) {
	sql :="select _id,title,ctime from simba_online limit 12000 export JSON 'test_export_json_1.txt' "
	RunAndPrintTheResult(sql)
}

// use for vector a-knn by OpenSearch
func TestKnnQuery(t *testing.T){
	sql := "select *,_id,_score from cv_knn_v1 where knn(field='embedding',k=10,vector=ARRAY(44.16973876953125,55.9416389465332,-1.7372,19.967805862426758,-4.899868965148926,-28.84906005859375,10.220306396484375,-30.81686782836914,-41.4800910949707,-15.460243225097656,60.88936996459961,10.749192237854004,-39.03456115722656,27.992897033691406,0.5167019963264465,32.380043029785156,5.034614562988281,-15.938974380493164,23.919647216796875,9.352738380432129,6.047122955322266,26.730600357055664,35.7619743347168,8.296953201293945,-25.745620727539062,-32.85493087768555,-23.82927131652832,-52.4396858215332,-12.47620964050293,-31.062381744384766,-51.89823913574219,27.583866119384766,0.2795354425907135,-3.898052215576172,-36.067935943603516,8.104389190673828,16.13810920715332,-3.9156153202056885,-21.607954025268555,-31.465408325195312,21.949987411499023,63.101226806640625,3.3646774291992188,-36.13178634643555,-48.12810134887695,4.8083109855651855,91.01653289794922,-9.161678314208984,64.76493835449219,14.686233520507812,34.73862075805664,-37.96909713745117,-38.57183074951172,32.844051361083984,56.04409408569336,47.58161544799805,-3.0995285511016846,-2.058413505554199,-23.581912994384766,-25.247804641723633,7.124398708343506,-40.18925857543945,10.577170372009277,-2.9566304683685303,-40.467857360839844,7.434725284576416,16.23598861694336,16.634864807128906,-35.62395095825195,-3.076709032058716,5.472452640533447,46.590370178222656,-34.859439849853516,11.366242408752441,39.7512321472168,-32.29405212402344,3.1306238174438477,-20.73850440979004,38.5838508605957,-20.409353256225586,18.580810546875,7.201180458068848,-7.2724928855896,12.908051490783691,-22.045106887817383,-48.814537048339844,16.926151275634766,-17.302837371826172,-4.468443870544434,5.056460857391357,-35.8611946105957,19.08583641052246,-19.218303680419922,-16.952903747558594,-24.75442123413086,14.415932655334473,46.52892303466797,2.638831615447998,-7.005253791809082,-18.0963191986084,-21.857624053955078,-20.881488800048828,12.002959251403809,-12.04485034942627,26.53418731689453,-6.888240814208984,34.20393753051758,-6.42675256729126,0.050941791385412216,-33.51689147949219,-30.886465072631836,-16.050399780277438,66.83118438720703,3.0258941650390625,-11.364269256591797,10.704959869384766,73.22728729248047,-38.62272262573242,-38.6381721496582,-31.399974822998047,-29.957828521728516,28.363208770751953,12.734430313110352,-21.987472534179688,-25.800979614257812,-9.703461647033691,45.4826774597168,7.640254497528076)) limit 1";
	RunAndPrintTheResultInSomeCluster( sql, "elastic-sql-rest-knn.yml")
}