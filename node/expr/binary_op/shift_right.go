package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n ShiftRight) Name() string {
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
