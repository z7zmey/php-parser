package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Spaceship struct {
	BinaryOp
}

func NewSpaceship(variable node.Node, expression node.Node) node.Node {
	return Spaceship{
		BinaryOp{
			node.SimpleNode{Name: "BinarySpaceship", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
