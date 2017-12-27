package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Minus) Name() string {
	return "Minus"
}

type Minus struct {
	AssignOp
}

func NewMinus(variable node.Node, expression node.Node) node.Node {
	return Minus{
		AssignOp{
			"AssignMinus",
			variable,
			expression,
		},
	}
}

func (n Minus) Walk(v node.Visitor) {
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
