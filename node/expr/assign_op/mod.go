package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Mod) Name() string {
	return "Mod"
}

type Mod struct {
	AssignOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		AssignOp{
			"AssignMod",
			variable,
			expression,
		},
	}
}

func (n Mod) Walk(v node.Visitor) {
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
