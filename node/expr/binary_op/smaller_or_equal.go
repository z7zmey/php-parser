package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n SmallerOrEqual) Name() string {
	return "SmallerOrEqual"
}

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(variable node.Node, expression node.Node) node.Node {
	return SmallerOrEqual{
		BinaryOp{
			"BinarySmallerOrEqual",
			variable,
			expression,
		},
	}
}
