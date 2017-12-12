package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	BinaryOp
}

func NewMinus(variable  node.Node, expression node.Node) node.Node {
	return Minus{
		BinaryOp{
			node.SimpleNode{Name: "BinaryMinus", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
