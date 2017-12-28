package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseAnd) Name() string {
	return "BitwiseAnd"
}

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		AssignOp{
			"AssignBitwiseAnd",
			variable,
			expression,
		},
	}
}

func (n BitwiseAnd) Walk(v node.Visitor) {
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
