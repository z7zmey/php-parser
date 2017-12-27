package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	BinaryOp
}

func NewBitwiseOr(variable node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		BinaryOp{
			"BinaryBitwiseOr",
			variable,
			expression,
		},
	}
}

func (n BitwiseOr) Name() string {
	return "BitwiseOr"
}

func (n BitwiseOr) Walk(v node.Visitor) {
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
