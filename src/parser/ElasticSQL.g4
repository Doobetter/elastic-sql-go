grammar ElasticSQL;

elasticSQL:
	statements+=statement ('|'statements+=statement)*; // ELASTIC-SQL 有多个分析元素以'|'连接组成

statement: 
	queryStatement 
	| analysisStatement 
	| sparkStatement 
	| insertStatement 
	| updateStatement
	| batchUpdateStatement
	| deleteStatement
	| descStatement
	| aliasStatement
	| alterStatement
	| joinQueryAnalysisStatement
	| hive2Statement
	| jdbcStatement
	| fileLoadStatement
	| analyzeStatement
	| hanLPStatement
	;

queryStatement: 
	(USING front = strictIdentifier)?
	SELECT  
	('/*+' (SCROLL EQ keep=str)? (SCORE ((EQ score=booleanValue)|(GE minScore=number)))? '*/')? 
//	(fields += fieldIdentifier (','fields +=fieldIdentifier)*)
//	(','? hlightFields=highlight)?
//	(',' exprFields=scriptFields)?
//	(',' innerFields=innerHit)?
    selectItems+=selectItem (','? selectItems+=selectItem)*
	(
		(BY SCROLL_ID EQ scroll_id = QUOTASTR)? 
	    |
		(
			FROM indexes += indexIdentifier(',' indexes +=indexIdentifier)* 
			
			(WHERE whereExpression  customScoreExpr? )?

			(rescoreExpr)?				
			
			(collapseExpr)?
			
			(
				ORDER BY sorts+=sortItem (',' sorts+=sortItem)*	
				//SEARCH AFTER '(' sortOff ')'
			)?
			(TRACK_HIT track=LONG)?
			( LIMIT  (offset=LONG ',')? limit = LONG
				
				memSort?
				
			)?
			// scroll slice
			(
				SLICE  ('(' sliceMax=LONG 	( ','sliceField=fieldIdentifier	)?')')?
			)?
		)?
	)
	(MAP statName = strictIdentifier)?
	(exportStatement) ?
	
;

selectItem:
	fieldIdentifier
	|highlight
	|scriptField
	|innerHit
;

collapseExpr: 
 COLLAPSE  BY  field=fieldIdentifier
;


customScoreExpr:
	// score_script 自定义score
	(CUSTOM_SCORE scriptPhrase)   
;

rescoreExpr: 
    // 结果集二次查询
	RESCORE (WINDOWS EQ win=LONG)? (WEIGHT'('pre=DOUBLE ',' sec=DOUBLE')')?  whereExpression
;

// 内存二次排序
memSort: MEM_SORT BY  sorts+=sortItem (',' sorts+=sortItem)* (KEEP size= LONG)? ;



innerHit: 
	INNER_HIT '('props+=prop (','props+=prop)*')' AS as = strictIdentifier
;

scriptField:
    script = scriptPhrase  AS as = strictIdentifier
;

highlight:
 HIGHLIGHT '('fieldAs (','fieldAs)*')' (BY tag=QUOTASTR)?
;

fieldAs:
 field = fieldIdentifier (AS? as = strictIdentifier)? (BY tag=QUOTASTR)?
;

whereExpression:
	 ('/*+' EXPLAIN ? (SCORE EQ score=booleanValue) '*/')? (ALL AND?)? logicalExpr?
;

logicalExpr: // 逻辑表达式
   comparableExpression
   | '(' inner = logicalExpr ')'
   | left = logicalExpr operator = AND  right = logicalExpr 
   | left = logicalExpr operator = OR  right = logicalExpr
;
comparableExpression: //比较表达式
	(not=NOT)?
	(  
		tCmp=termCompare // 普通字段比较
		| btwCmp = btwCompare // between and 
		| funcCmp = functionalCompare // 函数比较
		| mathCmp = arithmeticExpressionCompare // 算术表达比较
	)
;

// 对某字段比较
termCompare:
	field = fieldIdentifier 
	(POWER boost=number)? // boost
	operator = comparisonOperator  
	value = param
;

btwCompare: 
	field = fieldIdentifier (not=NOT)? 
	(
	(gte=BETWEEN a=param lte=AND b=param)
	|
	( RANGE (gte='[' | gt='(')  a=param ',' b=param (lte=']'| lt=')'))
	)
;



// 算术表达式间的比较
arithmeticExpressionCompare:
first = arithmeticExpression op=comparisonOperator second = multiplyingExpression 
;
// 加减乘除算术表达式，支持字符串用+连接
arithmeticExpression
   : first = multiplyingExpression (rest += addition)*
;
addition:
	op=(MINUS|PLUS ) next = multiplyingExpression
;
multiplyingExpression:
    first = atom  (rest += multi)* // 不能使用单独的字段做script?
;
multi:
	op=(MUL | DIV)  next = atom
;
atom
   : field = fieldIdentifier
   | TIME_T '(' field = fieldIdentifier ')'
   | data = param // number/string/boolean
   | '(' inner = arithmeticExpression ')'
;


// 函数表示的条件
functionalCompare:
	termLevelFunction
	 | 
	fullLevelFunction
	 |
	scriptFunction
	 |
	joinFunction
	
;

joinFunction: 
	funcName = joinFunctionNames '(' relationName=strictIdentifier (ON joinCondition=whereExpression)? ')'
;
joinFunctionNames: 	HAS_PARENT|HAS_CHILD;


scriptFunction: 
  script = scriptPhrase
//  |
//  arithmetic = arithmeticExpression 
;

fullLevelFunction: 
	funcName = fullLevelFunctionNames '('props+=prop (','props+=prop)*')'
;
fullLevelFunctionNames:
	QUERY_STRING
	| MATCH
	| MULTI_MATCH
	| MATCH_PHRASE
	| MATCH_PHRASE_PREFIX
	| KNN
;



//  field funcName(params) or funcName(field,params)
termLevelFunction:
	( field =fieldIdentifier (POWER boost=number)?)? funcName = termLevelFunctionNames '('params +=param2 (',' params +=param2)*')'
	|
	funcName = termLevelFunctionNames  '(' field =fieldIdentifier (POWER boost=number)? (',' params +=param2)*')'
//	|
//	funcName = termLevelFunctionNames (not=NOT)? '(' field =fieldIdentifier (POWER boost=number)? ',' useField = fieldIdentifier')'
	|
	( field =fieldIdentifier (POWER boost=number)?)? funcName = termLevelFunctionNames '(' useField = fieldIdentifier ')'
 ;

termLevelFunctionNames:
	IN|OUT
	|HAS_ALL|HAS_ANY
	|EXIST|MISS
	|RLIKE|LIKE|NOT_LIKE
	|IDS
	|STARTS_WITH
	|LOCAL_FILE

;


prop:
	// IDENTIFIER('^'LONG)?
	k= strictIdentifier
	EQ
	v=param
 ;

param:
	//arrayValue|booleanValue|LONG|DOUBLE|QUOTASTR|MUL
	arrayValue|booleanValue|LONG|DOUBLE|str|MUL
;
// func(field,param)分不开
param2:
	arrayValue|booleanValue|LONG|DOUBLE|QUOTASTR|MUL
;
arrayValue:
	ARRAY paramValues
;


sortItem:
	 (
	 	(field = fieldIdentifier )
	 	|  
	 	(dType=(STRING|NUMBER))? (script = scriptPhrase )
	 )
	 (filter='('whereExpression')')? 
	 (ordering=(ASC|DESC))? 
	 ( md=(SUM|AVG|MAX|MIN))?
;

exportStatement: 
	EXPORT DISTINCT?
	(
		'(' fields += exportField (','fields +=exportField)* ')'
	)?
	( HEADER ('(' heads += str (',' heads += str)* ')')? )?
	(
		fileType = (JSON | CSV |EXCEL) (SEP EQ sep= QUOTASTR)?
	)?

	// 保存文件名
	fileName = str
;


exportField:
	fieldIdentifier|metricNames
;

analysisStatement: 
	(USING front = strictIdentifier)?
	//aggStatement(','aggStatement)*
	aggStatement(';'aggStatement)*
;
aggStatement:
	
	(WHERE COUNT? where= whereExpression)?
	
	metrics += metricAgg (','metrics += metricAgg)*
	
	(GROUP BY bucketAggList)?
	
	havingExpr ?
	
	(LIMIT limit = LONG)?
	
	(MAP statName = strictIdentifier)?
	
	(exportStatement) ?
		
;
//metricAggList:
//	  
//;
metricAgg:
	// 小心 metricAs 使用了缺省字段 如 sum avg等 
	(WITH  with= whereExpression)?
 	(
 		metricName = metricNames '(' ')'
    |	
       	metricName = metricNames '(' field=fieldIdentifier (',' metricParams)?')' 
 	|
 		metricName = metricNames expr= '(' metricParams')' 
 	)
 	
 	(AS? metircAs = exportField )?
 	//(metircAs= asIdentifier )?

;




metricNames :
 	SCORE|COUNT|AVG|AVG_DATE|SUM|MAX|MIN|DISTINCT|VALUE_COUNT|PERCENTILES|PERCENTILE_RANKS|PERCENT|DISTINCT_PERCENT|TOP_HITS|POST_EXPR
;






metricParams:
	( kvs += mkv (',' kvs += mkv)* (',' script = scriptPhrase )?)
	|
	(script = scriptPhrase )
	
    
;
// key-value pair
mkv:
	k=metricParamNames EQ v=param
;
metricParamNames:
	MISSING|'precision_threshold'|'percents'|KEYED|SIZE|ORDER|INCLUDE|FORMAT
;

scriptPhrase:
  (SCRIPT  ('/*+' (LANG EQ lang=str) '*/')?  '(' script = QUOTASTR (',' props+=prop)* ')')
  |
  (MATH ? arithmetic = arithmeticExpression) 
;

bucketAggList:
	buckets += bucketAgg |'('buckets+=bucketAgg(',' buckets+=bucketAgg)*')'
;
// 为了添加 bucketAs
bucketAgg:
 	bucket = bucketAggChoice (AS? bucketAs = exportField )?
 	;
bucketAggChoice:
	termsBucket
	|rangeBucket
	|dateRangeBucket
	|histogramBucket
	|dateHistogramBucket
	|filtersBucket
	|significantBucket
;




// term group
termsBucket:	
	
	(field = fieldIdentifier)? 
    
	( (INCLUDE'('include+=param (',' include+=param )*')')
		| (EXCLUDE'('exclude+=param (','exclude+=param )*')')
		| (DOC_COUNT minDocCount = LONG) 
		| (MISSING AS missing=param )
		| (script = scriptPhrase)
		| (TOP top = LONG) 
		| ( ORDER BY orderPath = pathIdentifier (order = ASC | DESC)? )
	) *
;



havingExpr: 
	HAVING '(' bucketPath = pathIdentifier ',' script = str ')'
;

// range group
rangeBucket:
	RANGE'('
		field =fieldIdentifier  
		(',' (MISSING EQ)? missing=param)? 
		',' rangeExpr ')' 
;
rangeExpr:
	  ranges += rangeUnit  (',' ranges+=rangeUnit )*
;
rangeUnit:
	'['from=rangeFromTo ',' to=rangeFromTo (',' key=str)? ']'
;
rangeFromTo:
	MINUS|number
;

// date_range group
dateRangeBucket:
	DATE_RANGE'('
		field =fieldIdentifier 
		(',' (FORMAT EQ)? format=str)?  
		(',' MISSING EQ missing=str)? 
		',' dateRangeExpr')'
;

dateRangeExpr:
	  ranges += dateRange  (',' ranges+=dateRange )*
;
dateRange:
	'['from=dateRangeFromTo ',' to=dateRangeFromTo (',' key=str)? ']'
;

dateRangeFromTo: 
	MINUS|LONG|str
;

histogramBucket: 
	HISTOGRAM '('
		field =fieldIdentifier 
		',' (INTERVAL EQ)? interval=param
		(',' MISSING EQ  missing=param)? 
		(',' OFFSET EQ offset=param) ?
		(',' ORDER EQ order=pathIdentifier ASC?)?
		')'
;

dateHistogramBucket: 
	DATE_HISTOGRAM'('
		field =fieldIdentifier 
		',' (INTERVAL EQ)? interval=str 
		(',' FORMAT EQ format=str) ?  
		(',' MISSING EQ missing=param)? 
		(',' OFFSET EQ offset=param)?
		(',' ORDER EQ order=pathIdentifier ASC?)?
		 ')'
;

significantBucket: 
	SIGNIFICANT '('
		field =fieldIdentifier
		(
			(',' SAMPLE EQ smp = LONG )
			| (',' DOC_COUNT minDocCount = LONG) 
			| (',' SCORE EQ scoreAlg = str) 
			| (',' EXCLUDE'('exclude+=param (','exclude+=param )*')')
			| (',' INCLUDE'('include+=param (','include+=param )*')')
			| (',' TOP  top = LONG) 
		)*
	')'
;



filtersBucket: 
    field =fieldIdentifier 
	(WHEN where+= whereExpression THEN key+=str)+
	END

;

sparkStatement: 
	SQL 
	// 多个文件结构 hint
	('/*+' files+=dataStruct (','files += dataStruct)* '*/')?
	// SQL 语句
	sql=QUOTASTR
	(LIMIT limit = LONG)?
	(MAP statName = strictIdentifier)?
	(exportStatement) ?
;
dataStruct:
	STRUCT'('tableName=strictIdentifier ',' fields+=exportField (','fields+=exportField)* ')' 
;

// 客户端JOIN
joinQueryAnalysisStatement: 
	SELECT (fields += fieldIdentifier (','fields +=fieldIdentifier)*)
	(',' POST_EXPR'('arithmetic = arithmeticExpression')' (AS postAs= strictIdentifier)?)?
	FROM indexes += indexIdentifier(',' indexes +=indexIdentifier)* (AS aAs= strictIdentifier )?
	(WHERE whereExpression)?
	MEM? (kind=(LEFT|ALL|RIGHT))? JOIN tableB = strictIdentifier (AS bAs = strictIdentifier )?
	ON first = fieldIdentifier EQ second = fieldIdentifier
	(
		ORDER BY sorts+=sortItem (',' sorts+=sortItem)*	
	)?
	(LIMIT limit = LONG)?
	// scroll slice
	
	(
		SLICE  ('(' sliceMax=LONG 	( ','sliceField=fieldIdentifier	)?')')?
	)?
	(exportStatement) ?
	(AGG aggs += aggStatement (';' aggs += aggStatement)*)?
	
	(MAP statName = strictIdentifier)?
;





// 插入一行或者多行记录，不支持数组、Nested
insertStatement: 
	(USING front = strictIdentifier)? 
	MQL ? // 支持mysql插入，方便prepare statement
	INSERT INTO (sink=strictIdentifier '.')? table= indexName 
	(
		columns = fieldList
		(valueCluase = valueList )? 
	)?  
	// 
	(ON UUID? PK idField=fieldIdentifier)?
	
	(MAP statName = strictIdentifier)?
;



fieldList:
   '('(fields += fieldDefine (','fields +=fieldDefine)*)')'
;

valueList:
    VALUES valuesList+=paramValues (',' valuesList+=paramValues )?
;

paramValues:
	'(' vs+=param (',' vs+=param)* ')'
;

// 更新数据
updateStatement: 
	(UPDATE|UPSERT) indexes += indexIdentifier(',' indexes +=indexIdentifier)* 
	(
		(script =  scriptPhrase)?
		
		(SET fields+=updateField (','fields+=updateField )* )?
	)
	(
		(WHERE whereExpression) 
		| (BY IDS_FILE fileName = str )
		| (BY ID EQ id = str)
		| (BY '(' items += str (',' items+=str)* ')')
		
	)
	(MAP statName = strictIdentifier)?
;

updateField:
	(
		(field = fieldIdentifier EQ value = param)
		|
		(dType=(MAP|LIST) '('field=fieldIdentifier',' values += param (',' values += param )* ')')
	)                                                                                             #updateReplaceField
	|
	(ADD dType=(MAP|LIST) '('field=fieldIdentifier',' values += param (',' values += param ) * ')')  #updateAddField
	|
	(REMOVE dType=(MAP|LIST)  '('field=fieldIdentifier',' values += param (',' values += param )* ')')  #updateRemoveField
;
REMOVE: 'REMOVE' ;


// 批量更新和插入
batchUpdateStatement: 
	(BATCH_UPDATE|BATCH_UPSERT)  
	 docs += str (',' docs+=str)* 
	(
		(script =  scriptPhrase)?	
		(SET (fields += fieldIdentifier EQ values += param (','fields +=fieldIdentifier EQ values += param )*) )?
		
	)
	(MAP statName = strictIdentifier)?

;

deleteStatement: 
	DELETE FROM indexes += indexIdentifier(',' indexes +=indexIdentifier)* 
	(
		(WHERE whereExpression)
		| (BY IDS_FILE fileName = str )
		| (BY ID EQ id = str)
		| (BY IDS '(' ids = str')')
	)
	(MAP statName = strictIdentifier)?
;

descStatement: 
	(DESC|DESCRIBE) TEMPLATE? indexes += indexIdentifier(',' indexes +=indexIdentifier)* 
;

aliasStatement:
	(op=ADD ALIAS  aliasName= indexIdentifier TO indexes += indexIdentifier(',' indexes +=indexIdentifier)*) #addAlias
	|
	(op=DELETE ALIAS  aliasName= indexIdentifier FROM indexes += indexIdentifier(',' indexes +=indexIdentifier)*) #deleteAlias
	
;

alterStatement: 
	ALTER  indexes += indexIdentifier(',' indexes +=indexIdentifier)*  SET props+=prop (','props+=prop)*   
;



hive2Statement:
	(USING file = str)? 
	HQL sqls+=QUOTASTR (';' sqls+=QUOTASTR)*
	(MAP statName = strictIdentifier)?
	(exportStatement) ?
 ;

// 支持 CLICKHOUSE PRESTO MYSQL
jdbcStatement: 
	resourceType = (CK|PQL|MQL) 
	('/*+' URL EQ url=QUOTASTR 
		(USER EQ user=QUOTASTR )? 
		(PASSWORD EQ password=QUOTASTR)?  '*/' )?  
	subSQLs += basicSQL (';' subSQLs += basicSQL)*
;



basicSQL:
	sql = QUOTASTR (MAP statName = strictIdentifier) ? (exportStatement) ?
;

fileLoadStatement: 
	LOAD ALL?
	(
		DATA'(' fields += fieldDefine? (','fields += fieldDefine )*')' 
		|
		JSON dType=(MAP|LIST)?
	)
	FROM LOCAL? filePath=str 
	(SEPARATED BY sep=QUOTASTR)? // default ','
	(LIMIT  (offset=LONG ',')? limit = LONG)?
	(MAP statName = strictIdentifier)?
;



fieldDefine:
	fieldName =fieldIdentifier (fieldType =dataType)?
;

analyzeStatement: 
	ANALYZE content= QUOTASTR (BY analyzer= str (',' indexAnalyzer=str)? )?
;


hanLPStatement: 
	NLP  funcName=nlpFunc obj=str ENABLE '('enables+=str (',' enables+=str)* ')'
;

nlpFunc: SEGMENT|NAME|ORG|LOC;

dataType:
   BOOLEAN_T | STRING | LONG_T | DOUBLE_T | FLOAT_T | INTEGER_T |DATE_T | TIME_T | DATETIME_T | TIMESTAMP_T
;



// tokens start
comparisonOperator:
	LT | LE | GT | GE | NE | EQ
;

pathIdentifier:
	exportField (('.'|'>')exportField)* 
;
fieldIdentifier: 	
	(strictIdentifier (('.'|'$')strictIdentifier)*) | MUL
;
indexIdentifier:
	index=indexName ('.' indexType=strictIdentifier)? 
;
// 支持 index_test_* 通配符'*'
indexName:
	(IDENTIFIER MUL?)
	|
	QUOTASTR
;

strictIdentifier: IDENTIFIER|nonReserved;

str:
	QUOTASTR|IDENTIFIER|EXT_IDENTIFIER
;

number:
	num = (LONG|DOUBLE)
;
booleanValue
    : TRUE | FALSE
    ;
    
// nonReserved can only contain tokens
nonReserved
	:
	NUMBER |STRING|
	GROUP | SCORE | STRUCT | NOT | INTO | TO | ID | SLICE | JSON | CSV |SEP | SCRIPT | RETURN | INNER_HIT|
	//functions of where
	RLIKE | LIKE | IN | OUT | EXIST | MISS | QUERY_STRING | HAS_CHILD | HAS_PARENT | MATCH| MULTI_MATCH |MATH|MATCH_PHRASE_PREFIX|MATCH_PHRASE|KNN|SIZE|
	//functions of metric
	COUNT | MAX | MIN | AVG | SUM | DISTINCT | VALUE_COUNT | PERCENTILE_RANKS | PERCENTILES | MATH | TOP_HITS| PERCENT| DISTINCT_PERCENT|
	INTERVAL | FORMAT | KEYED | MISSING | NULL | END | RANGE| 
	// script
	SCRIPT| LANG |
	
	RESCORE |WINDOWS| COLLAPSE|
	// agg
	EXCLUDE | INCLUDE| TOP_HITS | SAMPLE|
	// ES non-query
	IDS | IDS_FILE | LOCAL| DATA| SEPARATED | ANALYZE| ADD |PUT| TEMPLATE| LIST| ALIAS| PK | UUID |
	// non-es
	SQL|CK|URL|PERCENT|URL|PASSWORD|USER |DATE_T | TIME_T | DATETIME_T | TIMESTAMP_T|SEP|EXCEL|JSON|CSV|FORMAT|
	MEM| LEFT| RIGHT |ALL|
	// NLP
	ORG | NLP | ENABLE | SEGMENT| NAME |LOC
	;

SELECT: 'SELECT';
SCROLL: 'SCROLL';
HIGHLIGHT: 'HIGHLIGHT';
SCROLL_ID: 'SCROLL_ID';
INNER_HIT: 'INNER_HIT';
WHERE: 'WHERE';
CUSTOM_SCORE: 'CUSTOM_SCORE' ;
RESCORE: 'RESCORE';
WINDOWS: 'WINDOWS';
COLLAPSE: 'COLLAPSE';
SCORE: 'SCORE';
MEM_SORT: 'MEM-SORT' ;
KEEP: 'KEEP';
WEIGHT: 'WEIGHT';
EXPLAIN: 'EXPLAIN' ;
WHERES: 'WHERES';
LIMIT: 'LIMIT';
TRACK_HIT: 'TRACK_HIT';
OFFSET: 'OFFSET';
SQL: 'SQL';
STRUCT: 'STRUCT';
HQL: 'HQL';
CK: 'CK';
PQL: 'PQL';
MQL: 'MQL';
URL: 'URL';
PASSWORD: 'PASSWORD';
USER: 'USER' ;

WITH: 'WITH';
AS: 'AS';
MAP: 'MAP';
USING: 'USING';
GROUP: 'GROUP';
BY: 'BY';
ORDER: 'ORDER';
ASC: 'ASC';
DESC: 'DESC';
DESCRIBE: 'DESCRIBE';
TEMPLATE: 'TEMPLATE' ;
BETWEEN: 'BETWEEN';
OR: 'OR';
ALL: 'ALL';
AND: 'AND';
NOT: 'NOT';
INTO: 'INTO';
TO: 'TO';
FROM: 'FROM';
HAVING: 'HAVING';
EXCLUDE: 'EXCLUDE';
INCLUDE: 'INCLUDE';
DOC_COUNT: 'DOC_COUNT';
TOP: 'TOP';
INSERT: 'INSERT';
VALUES: 'VALUES' ;
ID: 'ID';
SLICE: 'SLICE';
EXPORT: 'EXPORT';
HEADER: 'HEADER';
JSON: 'JSON';
CSV: 'CSV';
EXCEL: 'EXCEL';
SEP: 'SEP';
SCRIPT: 'SCRIPT'|'SCRIPT_SET';
LANG: 'LANG';
RETURN: 'RETURN';
ON: 'ON';
PK: 'PK';
UUID: 'UUID';
MEM: 'MEM';
RIGHT: 'RIGHT';
LEFT: 'LEFT';
JOIN: 'JOIN';
AGG: 'AGG';
ALIAS: 'ALIAS';
ALTER: 'ALTER';
//functions of where
RLIKE: 'RLIKE';
LIKE: 'LIKE';
NOT_LIKE: 'NOT_LIKE';
STARTS_WITH: 'STARTS_WITH';
IN: 'IN';
OUT: 'OUT';
HAS_ANY: 'HAS_ANY';
HAS_ALL: 'HAS_ALL';
EXIST: 'EXIST';
MISS: 'MISS';
QUERY_STRING: 'QUERY_STRING';
MATCH: 'MATCH';
MULTI_MATCH: 'MULTI_MATCH';
MATCH_PHRASE_PREFIX: 'MATCH_PHRASE_PREFIX';
MATCH_PHRASE: 'MATCH_PHRASE';
KNN: 'KNN';
LOCAL_FILE: 'LOCAL_FILE';
HAS_CHILD: 'HAS_CHILD';
HAS_PARENT: 'HAS_PARENT';

//functions of metric
COUNT: 'COUNT';
MAX: 'MAX';
MIN: 'MIN';
AVG: 'AVG';
AVG_DATE: 'AVG_DATE';
SUM: 'SUM';
DISTINCT: 'DISTINCT';
VALUE_COUNT: 'VALUE_COUNT';
PERCENTILE_RANKS: 'PERCENTILE_RANKS';
PERCENTILES: 'PERCENTILES';
PERCENT: 'PERCENT';
DISTINCT_PERCENT: 'DISTINCT_PERCENT';
TOP_HITS: 'TOP_HITS' ;
POST_EXPR: 'POST_EXPR' ;
MATH: 'MATH';

RANGE: 'RANGE';
DATE_RANGE: 'DATE_RANGE';
HISTOGRAM: 'HISTOGRAM';
DATE_HISTOGRAM: 'DATE_HISTOGRAM';
INTERVAL: 'INTERVAL';
FORMAT: 'FORMAT';
KEYED: 'KEYED';
SIZE: 'SIZE' ;
MISSING: 'MISSING';
NULL: 'NULL';
WHEN: 'WHEN';
END: 'END' ;
THEN: 'THEN';
SAMPLE: 'SAMPLE' ;
SIGNIFICANT: 'SIGNIFICANT';

IDS: 'IDS';
IDS_FILE: 'IDS_FILE';
SET: 'SET';
UPDATE: 'UPDATE';
UPSERT: 'UPSERT';
PUT: 'PUT';
ADD: 'ADD';
BATCH_UPDATE: 'BATCH_UPDATE' ;
BATCH_UPSERT: 'BATCH_UPSERT';
DELETE: 'DELETE';

ARRAY: 'ARRAY' ;
FALSE: 'FALSE';
TRUE: 'TRUE';
NUMBER: 'NUMBER';
STRING: 'STRING';

// data type
BOOLEAN_T: 'BOOLEAN';
LONG_T: 'LONG';
DOUBLE_T: 'DOUBLE';
INTEGER_T: 'INTEGER'|'INT';
DATETIME_T:'DATETIME';
TIMESTAMP_T: 'TIMESTAMP';
DATE_T: 'DATE';
TIME_T: 'TIME';
FLOAT_T: 'FLOAT';
// load data
LOCAL: 'LOCAL';
DATA: 'DATA';
LOAD: 'LOAD';
LIST: 'LIST';
SEPARATED: 'SEPARATED';
ANALYZE: 'ANALYZE';

NLP: 'NLP';
ENABLE: 'ENABLE';
SEGMENT: 'SEG'|'SEGMENT';
NAME: 'NAME';
ORG: 'ORG'|'ORGANIZATION';
LOC: 'LOC'| 'LOCATION';


// NON_LIMIT: MINUS;


PLUS:'+';

MINUS: '-' ;

MUL: '*' ;

DIV:'/';

MOD: '%' ;

POWER: '^';

DOT: '.';

//COMMA:',';

//relations   

EQ: '='|'==';

NE: '!=';

LT: '<' ;

LE: '<=' ;

GT: '>' ;

GE: '>=' ;

LONG:
	(SIGN)?DIGIT+
;
DOUBLE:
	//(SIGN)?DIGIT+ ('.' DIGIT+)?
(SIGN)?
    (
    DIGIT+ DOT DIGIT*
    | DOT DIGIT+
    | DIGIT+ (DOT DIGIT*)? EXPONENT
    | DOT DIGIT+ EXPONENT
    )
    ;

QUOTASTR: 
//'\'' ( ~('\''|'\\') | ('\\' .) )* '\''
 //| '"' ( ~('"'|'\\') | ('\\' .) )* '"'
 	'"' ( ~'"' | '""' )* '"'
 	|'\'' ( ~'\'' | '\'\'' )* '\''
    | '`' ( ~'`' | '``' )* '`'
;

IDENTIFIER: 
	(LETTER |'_' )(LETTER | DIGIT | '_')* 
;
EXT_IDENTIFIER:
  // digit head or contain '-'
	(DIGIT|LETTER |'_' )(LETTER | DIGIT | '_' | '-')*
;
fragment EXPONENT
    : 'E' SIGN? DIGIT+
    ;
// 包括中文
fragment LETTER
: [a-zA-Z\u4e00-\u9fcb];
fragment DIGIT
: [0-9];



fragment SIGN
: PLUS|MINUS
   ;

LINE_COMMENT : ('--' .*?'\n') -> skip  ;
BLOCK_COMMNET : '/*'~('+').*? '*/' -> skip ; // Match "/*" stuff "*/"

WS
: [ \t\r\n]+ -> skip
;