package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Plus) Name() string {
	return "Plus"
}

type Plus struct {
	BinaryOp
}

func NewPlus(variable node.Node, expression node.Node) node.Node {
	return Plus{
		BinaryOp{
			"BinaryPlus",
			variable,
			expression,
		},
	}
}
