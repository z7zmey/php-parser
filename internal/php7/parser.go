package php7

import (
	"bytes"

	"github.com/z7zmey/php-parser/internal/positionbuilder"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/token"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	Lexer           scanner.Scanner
	currentToken    *scanner.Token
	positionBuilder *positionbuilder.PositionBuilder
	rootNode        ast.Vertex
	errors          []*errors.Error
}

// NewParser creates and returns new Parser
func NewParser(src []byte, v string) *Parser {
	parser := &Parser{}

	lexer := scanner.NewLexer(src, func(e *errors.Error) {
		parser.errors = append(parser.errors, e)
	})
	lexer.PHPVersion = v

	parser.Lexer = lexer

	return parser
}

func (l *Parser) Lex(lval *yySymType) int {
	t := l.Lexer.Lex()

	l.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (l *Parser) Error(msg string) {
	var pos = l.currentToken.Position

	l.errors = append(l.errors, errors.NewError(msg, &pos))
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.errors
}

func (l *Parser) WithTokens() {
	l.Lexer.SetWithHiddenTokens(true)
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.errors = nil
	l.rootNode = nil
	l.positionBuilder = &positionbuilder.PositionBuilder{}

	// parse

	return yyParse(l)
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() ast.Vertex {
	return l.rootNode
}

// helpers

func lastNode(nn []ast.Vertex) ast.Vertex {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}

func firstNode(nn []ast.Vertex) ast.Vertex {
	return nn[0]
}

func isDollar(r rune) bool {
	return r == '$'
}

func (l *Parser) MoveFreeFloating(src ast.Vertex, dst ast.Vertex) {
	if l.Lexer.GetWithHiddenTokens() == false {
		return
	}

	if src.GetNode().Tokens == nil {
		return
	}

	l.setFreeFloating(dst, token.Start, src.GetNode().Tokens[token.Start])
	delete(src.GetNode().Tokens, token.Start)
}

func (l *Parser) setFreeFloating(dst ast.Vertex, p token.Position, strings []token.Token) {
	if l.Lexer.GetWithHiddenTokens() == false {
		return
	}

	if len(strings) == 0 {
		return
	}

	dstCollection := &dst.GetNode().Tokens
	if *dstCollection == nil {
		*dstCollection = make(token.Collection)
	}

	(*dstCollection)[p] = strings
}

func (l *Parser) GetFreeFloatingToken(t *scanner.Token) []token.Token {
	if l.Lexer.GetWithHiddenTokens() == false {
		return []token.Token{}
	}

	return []token.Token{
		{
			ID:    token.ID(t.ID),
			Value: t.Value,
		},
	}
}

func (l *Parser) addDollarToken(v ast.Vertex) {
	if l.Lexer.GetWithHiddenTokens() == false {
		return
	}

	l.setFreeFloating(v, token.Dollar, []token.Token{
		{
			ID:    token.ID('$'),
			Value: []byte("$"),
		},
	})
}

func (l *Parser) splitSemiColonAndPhpCloseTag(htmlNode ast.Vertex, prevNode ast.Vertex) {
	if l.Lexer.GetWithHiddenTokens() == false {
		return
	}

	semiColon := prevNode.GetNode().Tokens[token.SemiColon]
	delete(prevNode.GetNode().Tokens, token.SemiColon)
	if len(semiColon) == 0 {
		return
	}

	if semiColon[0].Value[0] == ';' {
		l.setFreeFloating(prevNode, token.SemiColon, []token.Token{
			{
				ID:    token.ID(';'),
				Value: semiColon[0].Value[0:1],
			},
		})
	}

	vlen := len(semiColon[0].Value)
	tlen := 2
	if bytes.HasSuffix(semiColon[0].Value, []byte("?>\n")) {
		tlen = 3
	}

	phpCloseTag := []token.Token{}
	if vlen-tlen > 1 {
		phpCloseTag = append(phpCloseTag, token.Token{
			ID:    token.T_WHITESPACE,
			Value: semiColon[0].Value[1 : vlen-tlen],
		})
	}

	phpCloseTag = append(phpCloseTag, token.Token{
		ID:    T_CLOSE_TAG,
		Value: semiColon[0].Value[vlen-tlen:],
	})

	l.setFreeFloating(htmlNode, token.Start, append(phpCloseTag, htmlNode.GetNode().Tokens[token.Start]...))
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
