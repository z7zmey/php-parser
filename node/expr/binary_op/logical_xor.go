package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n LogicalXor) Name() string {
	return "LogicalXor"
}

type LogicalXor struct {
	BinaryOp
}

func NewLogicalXor(variable node.Node, expression node.Node) node.Node {
	return LogicalXor{
		BinaryOp{
			"BinaryLogicalXor",
			variable,
			expression,
		},
	}
}
