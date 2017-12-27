package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Smaller) Name() string {
	return "Smaller"
}

type Smaller struct {
	BinaryOp
}

func NewSmaller(variable node.Node, expression node.Node) node.Node {
	return Smaller{
		BinaryOp{
			"BinarySmaller",
			variable,
			expression,
		},
	}
}
