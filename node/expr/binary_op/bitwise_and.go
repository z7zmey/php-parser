package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	BinaryOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		BinaryOp{
			node.SimpleNode{Name: "BinaryBitwiseAnd", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
