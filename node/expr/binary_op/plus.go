package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Plus) Name() string {
	return "Plus"
}

type Plus struct {
	BinaryOp
}

func NewPlus(variable node.Node, expression node.Node) node.Node {
	return Plus{
		BinaryOp{
			"BinaryPlus",
			variable,
			expression,
		},
	}
}

func (n Plus) Walk(v node.Visitor) {
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
