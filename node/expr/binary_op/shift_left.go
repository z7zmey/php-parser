package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShiftLeft) Name() string {
	return "ShiftLeft"
}

type ShiftLeft struct {
	BinaryOp
}

func NewShiftLeft(variable node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		BinaryOp{
			"BinaryShiftLeft",
			variable,
			expression,
		},
	}
}

func (n ShiftLeft) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.left != nil {
		vv := v.Children("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.Children("right")
		n.right.Walk(vv)
	}
}
