package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	AssignOp
}

func NewPow(Variable node.Node, Expression node.Node) *Pow {
	return &Pow{
		AssignOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *Pow) Attributes() map[string]interface{} {
	return nil
}

func (n *Pow) Position() *node.Position {
	return n.position
}

func (n *Pow) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Pow) Walk(v node.Visitor) {
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
