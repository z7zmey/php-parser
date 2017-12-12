package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	AssignOp
}

func NewConcat(variable  node.Node, expression node.Node) node.Node {
	return Concat{
		AssignOp{
			node.SimpleNode{Name: "AssignConcat", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
