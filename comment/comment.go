package comment

import "github.com/z7zmey/php-parser/node"

type Comment interface {
	String() string
}

type Comments map[node.Node][]Comment

func (c Comments) AddComments(node node.Node, comments []Comment) {
	c[node] = append(c[node], comments...)
}
