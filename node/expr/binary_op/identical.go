package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Identical) Name() string {
	return "Identical"
}

type Identical struct {
	BinaryOp
}

func NewIdentical(variable node.Node, expression node.Node) node.Node {
	return Identical{
		BinaryOp{
			"BinaryIdentical",
			variable,
			expression,
		},
	}
}
