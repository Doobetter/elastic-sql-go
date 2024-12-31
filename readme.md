elastic-sql-go使用类SQL查询、检索Elasticsearch、OpenSearch.

Thanks to antlr4/runtime/Go/antlr、olivere/elastic ...

## 1.basic use

```
sql :="select _id,nick,ctime from indexname"
elasticSQL:=ElasticSQL(sql,"")
elasticSQL.Init()
elasticSQL.Execute()
fmt.Println(common.JSONStr(elasticSQL.GetTheResultSet()))
```

## 2.multi-cluster support
在工程目录的conf文件夹下放置多个集群的配置文件
```
|--conf
  |--elastic-sql-rest-default.yml
  |--elastic-sql-rest-x.yml
```

代码中传入配置文件名即可（不是路径）
```
sql1:="select * from tablename"
elasticSQL:=ElasticSQL(sql) // 默认default配置文件
elasticSQL.Init()
elasticSQL.Execute()

// elastic-sql-rest-x.yml x根据集群集体情况定义
sql2:="select * from tablename*"
elasticSQL=ElasticSQL(sql, "elastic-sql-rest-x.yml") 
elasticSQL.Init()
elasticSQL.Execute()
```

## 3. 支持的语法

### 3.0 SQL基本结构
根据DSL的结构分为query部分、agg部分，通过|连接，可以有多个sql语句
```sql
select field,... from table1 
    where condition
    sort 
    limit
    export
|
metric group by field ...
|
select * from table2 where ... 

```

### 3.1 Query
> 基本的检索查询

- term
- terms
- range
- like
- match
- query_string
- match_phrase
- nested

todo 
- parent-child(join)


### 3.2 Aggregation
> 统计分析

// todo

### 3.3 导出文件
// todo

导出csv文件
```
select *,_id from index_name limit 10
export csv 'filename.txt'
```
导出json文件
```
select *,_id from index_name limit 10
export json 'filename.txt'
```

