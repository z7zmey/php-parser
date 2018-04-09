// Package php5 parses PHP5
package php5

import (
	"io"

	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/token"
)

var rootnode node.Node
var comments comment.Comments
var positions position.Positions
var positionBuilder position.Builder

var parentNode node.Node

// Parse the php5 parser entrypoint
func Parse(src io.Reader, fName string) (node.Node, comment.Comments, position.Positions, []*errors.Error) {
	yyDebug = 0
	yyErrorVerbose = true
	rootnode = stmt.NewStmtList([]node.Node{}) //reset
	comments = comment.Comments{}
	positions = position.Positions{}
	positionBuilder = position.Builder{Positions: &positions}

	lexer := newLexer(src, fName)
	yyParse(lexer)
	return rootnode, comments, positions, lexer.errors
}

// ListGetFirstNodeComments returns comments of a first node in the list
func ListGetFirstNodeComments(list []node.Node) []comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return comments[node]
}

type foreachVariable struct {
	node  node.Node
	byRef bool
}

type nodesWithEndToken struct {
	nodes    []node.Node
	endToken token.Token
}

type boolWithToken struct {
	value bool
	token *token.Token
}

type simpleIndirectReference struct {
	all  []*expr.Variable
	last *expr.Variable
}

type altSyntaxNode struct {
	node  node.Node
	isAlt bool
}
