package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n AssignRef) Name() string {
	return "AssignRef"
}

type AssignRef struct {
	AssignOp
}

func NewAssignRef(variable node.Node, expression node.Node) node.Node {
	return AssignRef{
		AssignOp{
			"AssignRef",
			variable,
			expression,
		},
	}
}

func (n AssignRef) Walk(v node.Visitor) {
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
