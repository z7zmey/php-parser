package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Minus) Name() string {
	return "Minus"
}

type Minus struct {
	BinaryOp
}

func NewMinus(variable node.Node, expression node.Node) node.Node {
	return Minus{
		BinaryOp{
			"BinaryMinus",
			variable,
			expression,
		},
	}
}
