package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftRight struct {
	BinaryOp
}

func NewShiftRight(variable  node.Node, expression node.Node) node.Node {
	return ShiftRight{
		BinaryOp{
			node.SimpleNode{Name: "BinaryShiftRight", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
