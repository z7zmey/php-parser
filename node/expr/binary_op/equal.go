package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Equal) Name() string {
	return "Equal"
}

type Equal struct {
	BinaryOp
}

func NewEqual(variable node.Node, expression node.Node) node.Node {
	return Equal{
		BinaryOp{
			"BinaryEqual",
			variable,
			expression,
		},
	}
}
