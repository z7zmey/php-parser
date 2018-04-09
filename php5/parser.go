package php5

import (
	"bufio"
	goToken "go/token"
	"io"

	"github.com/cznic/golex/lex"

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
	scanner.Lexer
	lastToken       *token.Token
	positionBuilder *position.Builder
	errors          []*errors.Error
	rootNode        node.Node
	comments        comment.Comments
	positions       position.Positions
}

// NewParser creates and returns new Parser
func NewParser(src io.Reader, fName string) *Parser {
	file := goToken.NewFileSet().AddFile(fName, -1, 1<<31-1)
	lx, err := lex.New(file, bufio.NewReader(src), lex.RuneClass(scanner.Rune2Class))
	if err != nil {
		panic(err)
	}

	scanner := scanner.Lexer{
		Lexer:         lx,
		StateStack:    []int{0},
		PhpDocComment: "",
		Comments:      nil,
	}

	return &Parser{
		scanner,
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
	yyDebug = 0
	yyErrorVerbose = true

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
