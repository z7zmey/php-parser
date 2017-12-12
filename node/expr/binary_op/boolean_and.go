package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanAnd struct {
	BinaryOp
}

func NewBooleanAnd(variable node.Node, expression node.Node) node.Node {
	return BooleanAnd{
		BinaryOp{
			node.SimpleNode{Name: "BinaryBooleanAnd", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
