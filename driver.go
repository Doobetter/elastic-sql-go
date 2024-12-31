package elasticsql

import (
	"errors"
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/parser"
	"github.com/antlr4-go/antlr/v4"
)

func ElasticSQL(mySQL, confFileName string) (*basic.ExeElasticSQLCtx, error) {

	input := parser.NewCaseInsensitiveStream(antlr.NewInputStream(mySQL))
	lexer := parser.NewElasticSQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	elasticSQLParser := parser.NewElasticSQLParser(stream)
	elasticSQLParser.AddParseListener(new(parser.ElasticSQLReactListener))
	tree := elasticSQLParser.ElasticSQL()
	//tree,err:= toTreeBySSL(elasticSQLParser)
	//if err != nil{
	//	//
	//	log.Println("ssl parse wrong",err)
	//	stream.Seek(0)
	//	elasticSQLParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLL)
	//	tree = elasticSQLParser.ExeElasticSQLCtx()
	//}

	var elasticSQL *basic.ExeElasticSQLCtx
	var err error
	if confFileName != "" {
		elasticSQL, err = basic.NewElasticSQLContextByConf(confFileName)
	} else {
		elasticSQL = basic.NewExeElasticSQLCtx()
	}
	if err != nil {
		return nil, err
	}
	elasticSQL.SQL = mySQL
	visitor := NewMyElasticVisitor(mySQL, elasticSQL)
	a := visitor.VisitElasticSQL(tree.(*parser.ElasticSQLContext))
	//a:=tree.Accept(visitor)
	if a != nil {
		return a.(*basic.ExeElasticSQLCtx), nil
	}
	return nil, errors.New("parse error no ElasticSQLContent return")
}
func toTreeBySSL(elasticSQLParser *parser.ElasticSQLParser) (antlr.ParserRuleContext, error) {
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
	return tree, err
}
