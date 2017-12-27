package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShiftLeft) Name() string {
	return "ShiftLeft"
}

type ShiftLeft struct {
	AssignOp
}

func NewShiftLeft(variable node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		AssignOp{
			"AssignShiftLeft",
			variable,
			expression,
		},
	}
}
