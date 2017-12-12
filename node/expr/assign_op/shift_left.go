package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftLeft struct {
	AssignOp
}

func NewShiftLeft(variable  node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		AssignOp{
			node.SimpleNode{Name: "AssignShiftLeft", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
