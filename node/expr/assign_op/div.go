package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	AssignOp
}

func NewDiv(variable node.Node, expression node.Node) node.Node {
	return &Div{
		AssignOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Div) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Div) Position() *node.Position {
	return n.position
}

func (n Div) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
