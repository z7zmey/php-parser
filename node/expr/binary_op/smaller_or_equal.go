package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(variable node.Node, expression node.Node) node.Node {
	return SmallerOrEqual{
		BinaryOp{
			node.SimpleNode{Name: "BinarySmallerOrEqual", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
