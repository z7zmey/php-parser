package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type GreaterOrEqual struct {
	BinaryOp
}

func NewGreaterOrEqual(variable node.Node, expression node.Node) node.Node {
	return GreaterOrEqual{
		BinaryOp{
			node.SimpleNode{Name: "BinaryGreaterOrEqual", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
