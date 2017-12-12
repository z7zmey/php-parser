package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Smaller struct {
	BinaryOp
}

func NewSmaller(variable node.Node, expression node.Node) node.Node {
	return Smaller{
		BinaryOp{
			node.SimpleNode{Name: "BinarySmaller", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
