package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftLeft struct {
	BinaryOp
}

func NewShiftLeft(variable  node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		BinaryOp{
			node.SimpleNode{Name: "BinaryShiftLeft", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
