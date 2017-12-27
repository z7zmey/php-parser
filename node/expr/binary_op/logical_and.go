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

func (n LogicalAnd) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.left != nil {
		vv := v.Children("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.Children("right")
		n.right.Walk(vv)
	}
}
