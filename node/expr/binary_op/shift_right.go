package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShiftRight) Name() string {
	return "ShiftRight"
}

type ShiftRight struct {
	BinaryOp
}

func NewShiftRight(variable node.Node, expression node.Node) node.Node {
	return ShiftRight{
		BinaryOp{
			"BinaryShiftRight",
			variable,
			expression,
		},
	}
}

func (n ShiftRight) Walk(v node.Visitor) {
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
