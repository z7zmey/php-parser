package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Mul) Name() string {
	return "Mul"
}

type Mul struct {
	BinaryOp
}

func NewMul(variable node.Node, expression node.Node) node.Node {
	return Mul{
		BinaryOp{
			"BinaryMul",
			variable,
			expression,
		},
	}
}
