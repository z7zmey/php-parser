package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n NotEqual) Name() string {
	return "NotEqual"
}

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(variable node.Node, expression node.Node) node.Node {
	return NotEqual{
		BinaryOp{
			"BinaryNotEqual",
			variable,
			expression,
		},
	}
}

func (n NotEqual) Walk(v node.Visitor) {
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
