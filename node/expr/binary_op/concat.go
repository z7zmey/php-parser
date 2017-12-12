package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	BinaryOp
}

func NewConcat(variable  node.Node, expression node.Node) node.Node {
	return Concat{
		BinaryOp{
			node.SimpleNode{Name: "BinaryConcat", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
