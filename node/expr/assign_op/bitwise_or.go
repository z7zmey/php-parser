package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseOr) Name() string {
	return "BitwiseOr"
}

type BitwiseOr struct {
	AssignOp
}

func NewBitwiseOr(variable node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		AssignOp{
			"AssignBitwiseOr",
			variable,
			expression,
		},
	}
}

func (n BitwiseOr) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.Children("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.Children("expression")
		n.expression.Walk(vv)
	}
}
