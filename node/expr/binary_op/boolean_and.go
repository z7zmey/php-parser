package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BooleanAnd) Name() string {
	return "BooleanAnd"
}

type BooleanAnd struct {
	BinaryOp
}

func NewBooleanAnd(variable node.Node, expression node.Node) node.Node {
	return BooleanAnd{
		BinaryOp{
			"BinaryBooleanAnd",
			variable,
			expression,
		},
	}
}

func (n BooleanAnd) Walk(v node.Visitor) {
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
