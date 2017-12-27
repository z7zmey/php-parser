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

func (n LogicalOr) Walk(v node.Visitor) {
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
