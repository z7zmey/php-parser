package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShiftRight) Name() string {
	return "ShiftRight"
}

type ShiftRight struct {
	AssignOp
}

func NewShiftRight(variable node.Node, expression node.Node) node.Node {
	return ShiftRight{
		AssignOp{
			"AssignShiftRight",
			variable,
			expression,
		},
	}
}

func (n ShiftRight) Walk(v node.Visitor) {
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
