package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Coalesce) Name() string {
	return "Coalesce"
}

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(variable node.Node, expression node.Node) node.Node {
	return Coalesce{
		BinaryOp{
			"BinaryCoalesce",
			variable,
			expression,
		},
	}
}

func (n Coalesce) Walk(v node.Visitor) {
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
