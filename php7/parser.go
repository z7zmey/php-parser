package php7

import (
	"io"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
	"github.com/z7zmey/php-parser/token"
)

func (lval *yySymType) Token(t token.Token) {
	lval.token = t
}

// Parser structure
type Parser struct {
	*scanner.Lexer
	path            string
	lastToken       *token.Token
	positionBuilder *position.Builder
	errors          []*errors.Error
	rootNode        node.Node
	comments        comment.Comments
	positions       position.Positions
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
	l.lastToken = &lval.token
	return t
}

func (l *Parser) Error(msg string) {
	l.errors = append(l.errors, errors.NewError(msg, *l.lastToken))
}

// Parse the php7 Parser entrypoint
func (l *Parser) Parse() int {
	// init
	l.errors = nil
	l.rootNode = nil
	l.comments = comment.Comments{}
	l.positions = position.Positions{}
	l.positionBuilder = &position.Builder{
		Positions: &l.positions,
	}

	// parse

	return yyParse(l)
}

func (l *Parser) listGetFirstNodeComments(list []node.Node) []comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return l.comments[node]
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
func (l *Parser) GetComments() comment.Comments {
	return l.comments
}

// GetPositions returns positions list
func (l *Parser) GetPositions() position.Positions {
	return l.positions
}
