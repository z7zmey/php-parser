package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Greater) Name() string {
	return "Greater"
}

type Greater struct {
	BinaryOp
}

func NewGreater(variable node.Node, expression node.Node) node.Node {
	return Greater{
		BinaryOp{
			"BinaryGreater",
			variable,
			expression,
		},
	}
}

func (n Greater) Walk(v node.Visitor) {
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
