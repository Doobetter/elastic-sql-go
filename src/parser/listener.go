package parser

import "github.com/antlr4-go/antlr/v4"

type ElasticSQLReactListener struct {
	BaseElasticSQLListener
}

func (l *ElasticSQLReactListener) exitNonReserved(c *ElasticSQLContext) {
	// replace nonReserved words with IDENT tokens
	c.GetParent().(antlr.ParserRuleContext).RemoveLastChild()

	token := c.GetChild(0).GetPayload().(antlr.Token)

	tkNew := antlr.NewCommonToken(token.GetSource(), ElasticSQLLexerIDENTIFIER, token.GetChannel(), token.GetStart(), token.GetStop())

	//tn := antlr.NewTerminalNodeImpl(tkNew)
	//c.GetParent().(antlr.ParserRuleContext).AddChild(tn)
	c.GetParent().(antlr.ParserRuleContext).AddTokenNode(tkNew)
}

type ParseErrorListener struct {
	antlr.DefaultErrorListener
	mySQL string
}
func  NewParseErrorListener(mySQL string) *ParseErrorListener{
	p:=new(ParseErrorListener)
	p.mySQL = mySQL
	return p
}
func (p *ParseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// TODO
}