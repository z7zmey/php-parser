package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Pow) Name() string {
	return "Pow"
}

type Pow struct {
	AssignOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		AssignOp{
			"AssignPow",
			variable,
			expression,
		},
	}
}

func (n Pow) Walk(v node.Visitor) {
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
