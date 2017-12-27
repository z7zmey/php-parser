package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Smaller) Name() string {
	return "Smaller"
}

type Smaller struct {
	BinaryOp
}

func NewSmaller(variable node.Node, expression node.Node) node.Node {
	return Smaller{
		BinaryOp{
			"BinarySmaller",
			variable,
			expression,
		},
	}
}

func (n Smaller) Walk(v node.Visitor) {
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
