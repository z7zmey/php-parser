package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	BinaryOp
}

func NewPow(variable  node.Node, expression node.Node) node.Node {
	return Pow{
		BinaryOp{
			node.SimpleNode{Name: "BinaryPow", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
