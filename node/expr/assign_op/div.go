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
