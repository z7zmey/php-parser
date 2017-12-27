package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n LogicalAnd) Name() string {
	return "LogicalAnd"
}

type LogicalAnd struct {
	BinaryOp
}

func NewLogicalAnd(variable node.Node, expression node.Node) node.Node {
	return LogicalAnd{
		BinaryOp{
			"BinaryLogicalAnd",
			variable,
			expression,
		},
	}
}
