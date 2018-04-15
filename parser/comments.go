package parser

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

// Comments a collection of comment groups assigned to nodes
type Comments map[node.Node][]*comment.Comment

// AddComments add comment group to the collection
func (c Comments) AddComments(node node.Node, comments []*comment.Comment) {
	c[node] = append(c[node], comments...)
}
