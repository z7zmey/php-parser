package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	BinaryOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		BinaryOp{
			"BinaryBitwiseXor",
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Name() string {
	return "BitwiseXor"
}

func (n BitwiseXor) Walk(v node.Visitor) {
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
