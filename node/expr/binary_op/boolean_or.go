package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BooleanOr) Name() string {
	return "BooleanOr"
}

type BooleanOr struct {
	BinaryOp
}

func NewBooleanOr(variable node.Node, expression node.Node) node.Node {
	return BooleanOr{
		BinaryOp{
			"BinaryBooleanOr",
			variable,
			expression,
		},
	}
}
