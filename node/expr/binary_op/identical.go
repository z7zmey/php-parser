package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Identical struct {
	BinaryOp
}

func NewIdentical(variable node.Node, expression node.Node) node.Node {
	return Identical{
		BinaryOp{
			node.SimpleNode{Name: "BinaryIdentical", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
