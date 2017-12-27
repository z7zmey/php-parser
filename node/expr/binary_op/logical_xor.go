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

func (n LogicalXor) Walk(v node.Visitor) {
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
