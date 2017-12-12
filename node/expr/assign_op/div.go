package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	AssignOp
}

func NewDiv(variable  node.Node, expression node.Node) node.Node {
	return Div{
		AssignOp{
			node.SimpleNode{Name: "AssignDiv", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
