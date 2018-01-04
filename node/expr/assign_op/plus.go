package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	AssignOp
}

func NewPlus(Variable node.Node, Expression node.Node) node.Node {
	return &Plus{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Plus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Plus) Position() *node.Position {
	return n.position
}

func (n Plus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Plus) Walk(v node.Visitor) {
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
