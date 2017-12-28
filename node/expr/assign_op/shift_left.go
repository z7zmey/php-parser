package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShiftLeft) Name() string {
	return "ShiftLeft"
}

type ShiftLeft struct {
	AssignOp
}

func NewShiftLeft(variable node.Node, expression node.Node) node.Node {
	return ShiftLeft{
		AssignOp{
			"AssignShiftLeft",
			variable,
			expression,
		},
	}
}

func (n ShiftLeft) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.GetChildrenVisitor("expression")
		n.expression.Walk(vv)
	}

	v.LeaveNode(n)
}
