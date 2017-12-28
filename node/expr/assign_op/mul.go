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
