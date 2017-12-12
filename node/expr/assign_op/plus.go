package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	AssignOp
}

func NewPlus(variable  node.Node, expression node.Node) node.Node {
	return Plus{
		AssignOp{
			node.SimpleNode{Name: "AssignPlus", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
