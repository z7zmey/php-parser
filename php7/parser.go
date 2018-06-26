package php7

import (
	"io"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/scanner"
)

func (lval *yySymType) Token(t *scanner.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	*scanner.Lexer
	path            string
	currentToken    *scanner.Token
	positionBuilder *parser.PositionBuilder
	errors          []*errors.Error
	rootNode        node.Node
	comments        parser.Comments
	positions       parser.Positions
}

// NewParser creates and returns new Parser
func NewParser(src io.Reader, path string) *Parser {
	lexer := scanner.NewLexer(src, path)

	return &Parser{
		lexer,
		path,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}
}

// Lex proxy to lexer Lex
func (l *Parser) Lex(lval *yySymType) int {
	t := l.Lexer.Lex(lval)
	l.currentToken = lval.token
	return t
}

func (l *Parser) Error(msg string) {
	l.errors = append(l.errors, errors.NewError(msg, l.currentToken))
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.errors = nil
	l.rootNode = nil
	l.comments = parser.Comments{}
	l.positions = parser.Positions{}
	l.positionBuilder = &parser.PositionBuilder{
		Positions:    &l.positions,
		PositionPool: &l.PositionPool,
	}

	// parse

	return yyParse(l)
}

// GetPath return path to file
func (l *Parser) GetPath() string {
	return l.path
}

// GetRootNode returns root node
func (l *Parser) GetRootNode() node.Node {
	return l.rootNode
}

// GetErrors returns errors list
func (l *Parser) GetErrors() []*errors.Error {
	return l.errors
}

// GetComments returns comments list
func (l *Parser) GetComments() parser.Comments {
	return l.comments
}

// GetPositions returns positions list
func (l *Parser) GetPositions() parser.Positions {
	return l.positions
}

// helpers

func lastNode(nn []node.Node) node.Node {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}

func firstNode(nn []node.Node) node.Node {
	return nn[0]
}

func isDollar(r rune) bool {
	return r == '$'
}

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.TokenPool.Put(yyDollar[i].token)
			p.PositionPool.Put(yyDollar[i].token.Position)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}
