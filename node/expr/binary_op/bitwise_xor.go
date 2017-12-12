package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	BinaryOp
}

func NewBitwiseXor(variable  node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		BinaryOp{
			node.SimpleNode{Name: "BinaryBitwiseXor", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
