package php5

import (
	"bytes"

	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/token"
)

// Parser structure
type Parser struct {
	Lexer          *scanner.Lexer
	currentToken   *scanner.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
}

// NewParser creates and returns new Parser
func NewParser(lexer *scanner.Lexer, errHandlerFunc func(*errors.Error)) *Parser {
	return &Parser{
		Lexer:          lexer,
		errHandlerFunc: errHandlerFunc,
	}
}

// Lex proxy to scanner Lex
func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	var pos = p.currentToken.Position
	p.errHandlerFunc(errors.NewError(msg, &pos))
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil
	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Vertex {
	return p.rootNode
}

// helpers

func lastNode(nn []ast.Vertex) ast.Vertex {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}

func (p *Parser) MoveFreeFloating(src ast.Vertex, dst ast.Vertex) {
	if _, ok := src.GetNode().Tokens[token.Start]; !ok {
		return
	}

	if src.GetNode().Tokens == nil {
		return
	}

	dstCollection := &dst.GetNode().Tokens
	if *dstCollection == nil {
		*dstCollection = make(token.Collection)
	}

	(*dstCollection)[token.Start] = src.GetNode().Tokens[token.Start]
	delete(src.GetNode().Tokens, token.Start)
}

func (p *Parser) setFreeFloating(dst ast.Vertex, pos token.Position, tokens []token.Token) {
	if len(tokens) == 0 {
		return
	}

	dstCollection := &dst.GetNode().Tokens
	if *dstCollection == nil {
		*dstCollection = make(token.Collection)
	}

	l := len(tokens)
	for _, v := range tokens[0 : l-1] {
		(*dstCollection)[pos] = append((*dstCollection)[pos], v)
	}
}

func (p *Parser) setFreeFloatingTokens(dst ast.Vertex, pos token.Position, tokens []token.Token) {
	if len(tokens) == 0 {
		return
	}

	dstCollection := &dst.GetNode().Tokens
	if *dstCollection == nil {
		*dstCollection = make(token.Collection)
	}

	(*dstCollection)[pos] = make([]token.Token, 0)

	for _, v := range tokens {
		(*dstCollection)[pos] = append((*dstCollection)[pos], v)
	}
}

func (p *Parser) setToken(dst ast.Vertex, pos token.Position, tokens []token.Token) {
	if len(tokens) == 0 {
		return
	}

	dstCollection := &dst.GetNode().Tokens
	if *dstCollection == nil {
		*dstCollection = make(token.Collection)
	}

	l := len(tokens)
	(*dstCollection)[pos] = append((*dstCollection)[pos], tokens[l-1])
}

func (p *Parser) splitSemiColonAndPhpCloseTag(htmlNode ast.Vertex, prevNode ast.Vertex) {
	if _, ok := prevNode.GetNode().Tokens[token.SemiColon]; !ok {
		return
	}

	semiColon := prevNode.GetNode().Tokens[token.SemiColon]
	delete(prevNode.GetNode().Tokens, token.SemiColon)
	if len(semiColon) == 0 {
		return
	}

	if semiColon[0].Value[0] == ';' {
		p.setFreeFloatingTokens(prevNode, token.SemiColon, []token.Token{
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

	p.setFreeFloatingTokens(htmlNode, token.Start, append(phpCloseTag, htmlNode.GetNode().Tokens[token.Start]...))
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
