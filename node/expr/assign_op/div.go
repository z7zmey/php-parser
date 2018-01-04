package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	AssignOp
}

func NewDiv(Variable node.Node, Expression node.Node) node.Node {
	return &Div{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
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

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expression != nil {
		vv := v.GetChildrenVisitor("Expression")
		n.Expression.Walk(vv)
	}

	v.LeaveNode(n)
}
