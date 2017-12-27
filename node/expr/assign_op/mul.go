package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Mul) Name() string {
	return "Mul"
}

type Mul struct {
	AssignOp
}

func NewMul(variable node.Node, expression node.Node) node.Node {
	return Mul{
		AssignOp{
			"AssignMul",
			variable,
			expression,
		},
	}
}

func (n Mul) Walk(v node.Visitor) {
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
