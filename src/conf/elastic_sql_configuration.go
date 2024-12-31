package conf

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var DefaultElasticSQLConf *ElasticSQLConfiguration

func init() {
	DefaultElasticSQLConf, _ = LoadAndNewElasticSQLConfiguration(DEFUALT_CONF_FILENAME)
}

const (
	DEFUALT_CONF_FILENAME = "elastic-sql-rest-default.yml"
)

type ElasticSQLConfiguration struct {
	ESClient struct {
		HttpHosts []string `yaml:"http_hosts"`
		Username  string   `yaml:"username"`
		Password  string   `yaml:"password"`
	} `yaml:"es_client"`
	DataDir string `yaml:"data_dir"`
	Query   struct {
		DefaultSize          int `yaml:"default_size"`
		ScrollSize           int `yaml:"scroll_size"`
		ScrollKeepTime       int `yaml:"scroll_keep_time"`
		ScrollSliceMax       int `yaml:"scroll_slice_max"`
		ScrollSliceThreshold int `yaml:"scroll_slice_threshold"`
	} `yaml:"query"`
	Agg struct {
		DefaultSize       int `yaml:"default_size"`
		DistinctPrecision int `yaml:"distinct_precision"`
	} `yaml:"agg"`
}

//ELASTIC_REST_CLIENT_HTTP_HOSTS []string  `yaml:"elastic.rest.client.http_hosts"`
//ELASTIC_REST_CLIENT_USERNAME string `yaml:"elastic.rest.client.username"`
//
//ELASTIC_REST_CLIENT_PASSWORD string `yml:"elastic.rest.client.password"`
//
//
//// query statement 默认返回结果条数
//ELASTIC_SQL_QUERY_DEFAULT_SIZE int `yml:"elastic.sql.query.default.size"`
//
//// scroll. 一次fetch获取记录数
//ELASTIC_SQL_FETCH_SCROLL_SIZE int `yml:"elastic.sql.fetch.scroll.size"`
//
//// scroll. how long it should keep the search basic alive
//ELASTIC_SQL_FETCH_SCROLL_KEEP_TIME  int `yml:"elastic.sql.fetch.scroll.keep.time"`
//
//// scroll. 最大slice,可以认为是获取数据的并行度
//ELASTIC_SQL_FETCH_SCROLL_SLICE_MAX   int `yml:"elastic.sql.fetch.scroll.slice.max"`
//
//// scroll. threshold to use slice, if total less than threshold use scroll, else use slice scroll
//ELASTIC_SQL_FETCH_SCROLL_SLICE_THRESHOLD   int `yml:"elastic.sql.fetch.scroll.slice.threshold"`
//
//// agg top default
//ELASTIC_SQL_AGG_GROUPING_DEFAULT_SIZE   int `yml:"elastic.sql.agg.grouping.default.size"`
//
//
////es-join 最大线程数，默认最小为2
//ELASTIC_SQL_JOIN_MULTI_TASK_QUERY_MAX_THREAD  int `yml:"elastic.sql.join.multi.task.query.max.thread"`
//
//ELASTIC_SQL_JOIN_MULTI_TASK_QUERY_QUEUE_SIZE  int `yml:"elastic.sql.join.multi.task.query.queue.size"`
//// 每个子任务的关联的条数，也是多线程和单线程执行的门限
//ELASTIC_SQL_JOIN_SUB_ROW_NUM  int `yml:"elastic.sql.join.sub.row.num"`

func LoadAndNewElasticSQLConfiguration(confFileName string) (*ElasticSQLConfiguration, error) {
	realName := confFileName
	// 当前目录
	if !PathExists(realName) {
		// conf文件夹下
		realName = "conf/" + confFileName
		if !PathExists(realName) {
			// ../conf文件夹下
			realName = "../conf/" + confFileName
			if !PathExists(realName) {
				realName = "/data1/search/conf/" + confFileName
				if !PathExists(realName) {
					return nil, errors.New("can't find conf file")
				}
			}
		}
	}
	bs, err := ioutil.ReadFile(realName)
	if err != nil {
		return nil, errors.New("can't read conf file")
	}
	conf := new(ElasticSQLConfiguration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		return nil, fmt.Errorf("can't parse conf file %s", confFileName)
	}
	return conf, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
