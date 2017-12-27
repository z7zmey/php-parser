package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Pow) Name() string {
	return "Pow"
}

type Pow struct {
	BinaryOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		BinaryOp{
			"BinaryPow",
			variable,
			expression,
		},
	}
}

func (n Pow) Walk(v node.Visitor) {
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
