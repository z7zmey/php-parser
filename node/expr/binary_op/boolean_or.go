package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanOr struct {
	BinaryOp
}

func NewBooleanOr(variable node.Node, expression node.Node) node.Node {
	return BooleanOr{
		BinaryOp{
			node.SimpleNode{Name: "BinaryBooleanOr", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
