package basic

import (
	"context"
	"errors"
	"github.com/Doobetter/elastic-sql-go/src/client"
	"github.com/Doobetter/elastic-sql-go/src/conf"
	"github.com/olivere/elastic/v7"
	"strconv"
	"sync"
)

var AllConfs = make(map[string]*conf.ElasticSQLConfiguration)
var AllESClientConns = make(map[string]*client.ESClientConnection)

var lock sync.Mutex

//ExeElasticSQLCtx 合并ElaticSQL和Context
type ExeElasticSQLCtx struct {
	GCtx        context.Context
	Conf        *conf.ElasticSQLConfiguration
	Conn        *client.ESClientConnection
	SQL         string
	Explain     string // DSL
	ExplainHead string // e.g. POST _search
	Canceled    bool
	_cnt        int // for Name num

	ProcessUnitsMap map[string]Statement // statementName -> Statement
	Statements      []Statement          // 保存stat顺序

	Results map[string]*ResultSet
}

func NewExeElasticSQLCtx() *ExeElasticSQLCtx {
	ctx := &ExeElasticSQLCtx{}
	ctx.GCtx = context.Background()
	ctx.Conf = conf.DefaultElasticSQLConf
	ctx.Conn = client.GetDefaultClientConnection()
	ctx.ProcessUnitsMap = make(map[string]Statement)
	ctx.Results = make(map[string]*ResultSet)
	return ctx
}

func NewElasticSQLContextByConf(confFileName string) (*ExeElasticSQLCtx,error) {
	ctx := new(ExeElasticSQLCtx)
	ctx.GCtx = context.Background()
	var err error
	ctx.Conn,err = GetESClientConnection(confFileName)
	if err!=nil{
		return nil,err
	}
	if ctx.Conn != nil{
		ctx.Conf = ctx.Conn.Conf
	}else {
		return ctx,errors.New("no client got")
	}
	ctx.ProcessUnitsMap = make(map[string]Statement)
	ctx.Results = make(map[string]*ResultSet)
	return ctx,nil
}

func GetESClientConnection(confFileName string) (*client.ESClientConnection,error) {
	if conn, ok := AllESClientConns[confFileName]; ok {
		return conn, nil
	} else {
		lock.Lock()
		defer lock.Unlock()
		conf, err := conf.LoadAndNewElasticSQLConfiguration(confFileName)
		if err != nil {
			return nil,err
		}
		conn,err:= client.NewESClientConnection(conf)
		if err!=nil{
			return nil,err
		}
		AllESClientConns[confFileName] = conn
		return conn,nil
	}
}

func getElasticSQLConfiguration(confFileName string) *conf.ElasticSQLConfiguration {
	config, ok := AllConfs[confFileName]
	if ok {
		return config
	} else {
		config, err := conf.LoadAndNewElasticSQLConfiguration(confFileName)
		if err != nil {
			// TODO
			return nil
		}

		return config
	}
}

func (c *ExeElasticSQLCtx) GetElasticClient() *elastic.Client {
	if c.Conn == nil {
		// 使用默认的
		return client.GetDefaultClientConnection().Client
	}
	return c.Conn.Client
}

func (c *ExeElasticSQLCtx) AddStatement(name string, statement Statement) {
	if name == "" {
		// set default statement Name.
		name = strconv.Itoa(c._cnt) + statement.DefaultNameSuffix()
		c._cnt++
		statement.SetName(name)
	}
	// TODO
	c.ProcessUnitsMap[name] = statement
	c.Statements = append(c.Statements, statement)
}

func (c *ExeElasticSQLCtx) GetResultSet(name string) *ResultSet {
	if rs, ok := c.Results[name]; ok {
		return rs
	} else {
		return nil
	}
}

//GetTheResultSet 获取唯一的result
func (c *ExeElasticSQLCtx) GetTheResultSet() *ResultSet {
	for _, v := range c.Results {
		return v
	}
	return nil
}

func (c *ExeElasticSQLCtx) AddResultSet(name string, rs *ResultSet) {
	c.Results[name] = rs
}

func (c *ExeElasticSQLCtx) Cancel() {
	if c.Canceled == false {
		c.Canceled = true
		// TODO
	}
}

func (c *ExeElasticSQLCtx) Init() {
	for _, stat := range c.Statements {
		stat.GenPostProcessCode()
		stat.Init(c)
	}
}
func (c *ExeElasticSQLCtx) Execute() {
	for _, stat := range c.Statements {
		stat.Execute(c)
	}
}
