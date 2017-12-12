package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalAnd struct {
	BinaryOp
}

func NewLogicalAnd(variable node.Node, expression node.Node) node.Node {
	return LogicalAnd{
		BinaryOp{
			node.SimpleNode{Name: "BinaryLogicalAnd", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
