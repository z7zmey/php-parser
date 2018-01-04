package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignRef struct {
	AssignOp
}

func NewAssignRef(Variable node.Node, Expression node.Node) node.Node {
	return &AssignRef{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n AssignRef) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AssignRef) Position() *node.Position {
	return n.position
}

func (n AssignRef) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AssignRef) Walk(v node.Visitor) {
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
