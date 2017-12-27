package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n LogicalOr) Name() string {
	return "LogicalOr"
}

type LogicalOr struct {
	BinaryOp
}

func NewLogicalOr(variable node.Node, expression node.Node) node.Node {
	return LogicalOr{
		BinaryOp{
			"BinaryLogicalOr",
			variable,
			expression,
		},
	}
}
