package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Spaceship) Name() string {
	return "Spaceship"
}

type Spaceship struct {
	BinaryOp
}

func NewSpaceship(variable node.Node, expression node.Node) node.Node {
	return Spaceship{
		BinaryOp{
			"BinarySpaceship",
			variable,
			expression,
		},
	}
}
