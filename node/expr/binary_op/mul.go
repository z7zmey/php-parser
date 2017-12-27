package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Mul) Name() string {
	return "Mul"
}

type Mul struct {
	BinaryOp
}

func NewMul(variable node.Node, expression node.Node) node.Node {
	return Mul{
		BinaryOp{
			"BinaryMul",
			variable,
			expression,
		},
	}
}

func (n Mul) Walk(v node.Visitor) {
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
