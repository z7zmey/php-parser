package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftRight struct {
	AssignOp
}

func NewShiftRight(variable  node.Node, expression node.Node) node.Node {
	return ShiftRight{
		AssignOp{
			node.SimpleNode{Name: "AssignShiftRight", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
