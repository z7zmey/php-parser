package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Pow) Name() string {
	return "Pow"
}

type Pow struct {
	BinaryOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		BinaryOp{
			"BinaryPow",
			variable,
			expression,
		},
	}
}
