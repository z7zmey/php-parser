package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(variable node.Node, expression node.Node) node.Node {
	return NotEqual{
		BinaryOp{
			node.SimpleNode{Name: "BinaryNotEqual", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
