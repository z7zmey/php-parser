package php5

import (
	"io"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/meta"
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

func (l *Parser) WithMeta() {
	l.Lexer.WithMeta = true
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.errors = nil
	l.rootNode = nil
	l.positionBuilder = &parser.PositionBuilder{}

	// parse

	return yyParse(l)
}

func (l *Parser) listGetFirstNodeMeta(list []node.Node) []meta.Meta {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return node.GetMeta()
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

func addMeta(n node.Node, mm []meta.Meta, tn meta.TokenName) {
	for _, m := range mm {
		m.SetTokenName(tn)
	}

	n.AddMeta(mm)
}

func (p *Parser) returnTokenToPool(yyDollar []yySymType, yyVAL *yySymType) {
	for i := 1; i < len(yyDollar); i++ {
		if yyDollar[i].token != nil {
			p.TokenPool.Put(yyDollar[i].token)
		}
		yyDollar[i].token = nil
	}
	yyVAL.token = nil
}
