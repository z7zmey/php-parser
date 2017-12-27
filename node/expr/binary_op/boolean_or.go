package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BooleanOr) Name() string {
	return "BooleanOr"
}

type BooleanOr struct {
	BinaryOp
}

func NewBooleanOr(variable node.Node, expression node.Node) node.Node {
	return BooleanOr{
		BinaryOp{
			"BinaryBooleanOr",
			variable,
			expression,
		},
	}
}

func (n BooleanOr) Walk(v node.Visitor) {
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
