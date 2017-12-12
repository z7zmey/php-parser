package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Equal struct {
	BinaryOp
}

func NewEqual(variable node.Node, expression node.Node) node.Node {
	return Equal{
		BinaryOp{
			node.SimpleNode{Name: "BinaryEqual", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
