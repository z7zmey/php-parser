package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Div) Name() string {
	return "Div"
}

type Div struct {
	AssignOp
}

func NewDiv(variable node.Node, expression node.Node) node.Node {
	return Div{
		AssignOp{
			"AssignDiv",
			variable,
			expression,
		},
	}
}

func (n Div) Walk(v node.Visitor) {
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
