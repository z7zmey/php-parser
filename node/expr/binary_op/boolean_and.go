package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BooleanAnd) Name() string {
	return "BooleanAnd"
}

type BooleanAnd struct {
	BinaryOp
}

func NewBooleanAnd(variable node.Node, expression node.Node) node.Node {
	return BooleanAnd{
		BinaryOp{
			"BinaryBooleanAnd",
			variable,
			expression,
		},
	}
}
