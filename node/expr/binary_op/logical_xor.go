package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalXor struct {
	BinaryOp
}

func NewLogicalXor(variable node.Node, expression node.Node) node.Node {
	return LogicalXor{
		BinaryOp{
			node.SimpleNode{Name: "BinaryLogicalXor", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
