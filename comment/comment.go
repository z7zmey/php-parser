package comment

import "github.com/z7zmey/php-parser/node"

// Comment represents comment lines in the code
type Comment interface {
	String() string
}

// Comments a collection of comment groups assigned to nodes
type Comments map[node.Node][]Comment

// AddComments add comment group to the collection
func (c Comments) AddComments(node node.Node, comments []Comment) {
	c[node] = append(c[node], comments...)
}
