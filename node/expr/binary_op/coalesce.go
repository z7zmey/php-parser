package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(variable node.Node, expression node.Node) node.Node {
	return Coalesce{
		BinaryOp{
			node.SimpleNode{Name: "BinaryCoalesce", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
