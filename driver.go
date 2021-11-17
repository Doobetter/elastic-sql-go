package elasticsql

import (
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ElasticSQL(mySQL, confFileName string ) *basic.ExeElasticSQLCtx {

	input:=parser.NewCaseInsensitiveStream(antlr.NewInputStream(mySQL))
	lexer:=parser.NewElasticSQLLexer(input)
	stream:=antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)
	elasticSQLParser:=parser.NewElasticSQLParser(stream)
	elasticSQLParser.AddParseListener(new(parser.ElasticSQLReactListener))
	tree:=elasticSQLParser.ElasticSQL()
	//tree,err:= toTreeBySSL(elasticSQLParser)
	//if err != nil{
	//	//
	//	log.Println("ssl parse wrong",err)
	//	stream.Seek(0)
	//	elasticSQLParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLL)
	//	tree = elasticSQLParser.ExeElasticSQLCtx()
	//}

	var elasticSQL *basic.ExeElasticSQLCtx
	if confFileName!=""{
		elasticSQL = basic.NewElasticSQLContextByConf(confFileName)
	}else {
		elasticSQL = basic.NewExeElasticSQLCtx()
	}
	elasticSQL.SQL = mySQL
	visitor := NewMyElasticVisitor(mySQL,elasticSQL)
	a := visitor.VisitElasticSQL(tree.(*parser.ElasticSQLContext))
	//a:=tree.Accept(visitor)
	if a !=nil{
		return a.(*basic.ExeElasticSQLCtx)
	}
	return nil
}
func toTreeBySSL(elasticSQLParser * parser.ElasticSQLParser) (antlr.ParserRuleContext,error){
	var tree antlr.ParserRuleContext
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
			err = r.(error)
			tree = nil
		}

	}()

	elasticSQLParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree = elasticSQLParser.ElasticSQL()
	return tree,err
}


