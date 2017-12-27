package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Assign) Name() string {
	return "Assign"
}

type Assign struct {
	AssignOp
}

func NewAssign(variable node.Node, expression node.Node) node.Node {
	return Assign{
		AssignOp{
			"Assign",
			variable,
			expression,
		},
	}
}

func (n Assign) Walk(v node.Visitor) {
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
