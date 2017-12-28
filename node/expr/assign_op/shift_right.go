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
