package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	BinaryOp
}

func NewBitwiseOr(variable  node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		BinaryOp{
			node.SimpleNode{Name: "BinaryBitwiseOr", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
