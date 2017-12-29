package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

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

func (n ShiftLeft) Name() string {
	return "ShiftLeft"
}

func (n ShiftLeft) Attributes() map[string]interface{} {
	return nil
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
