package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalOr struct {
	BinaryOp
}

func NewLogicalOr(variable node.Node, expression node.Node) node.Node {
	return LogicalOr{
		BinaryOp{
			node.SimpleNode{Name: "BinaryLogicalOr", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
