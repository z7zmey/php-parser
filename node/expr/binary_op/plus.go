package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	BinaryOp
}

func NewPlus(variable  node.Node, expression node.Node) node.Node {
	return Plus{
		BinaryOp{
			node.SimpleNode{Name: "BinaryPlus", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
