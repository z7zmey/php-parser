package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Greater struct {
	BinaryOp
}

func NewGreater(variable node.Node, expression node.Node) node.Node {
	return Greater{
		BinaryOp{
			node.SimpleNode{Name: "BinaryGreater", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
