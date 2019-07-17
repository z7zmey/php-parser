package php7

import (
	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/parser"
	"github.com/z7zmey/php-parser/internal/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	parser.AbstractParser
	yyParserImpl
}

// NewParser creates and returns new Parser
func NewParser() *Parser {
	return &Parser{
		parser.AbstractParser{
			Lexer: scanner.NewLexer(nil),
		},
		yyParserImpl{},
	}
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex(lval)
	p.CurrentToken = lval.token
	return t
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse(src []byte, a *graph.AST) int {
	p.AbstractParser.Reset(src, a)
	return p.yyParserImpl.Parse(p)
}

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.Lexer.ReturnTokenToPool(yyDollar[i].token)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}
