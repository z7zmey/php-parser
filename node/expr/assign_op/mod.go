package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	AssignOp
}

func NewMod(Variable node.Node, Expression node.Node) node.Node {
	return &Mod{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Mod) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Mod) Position() *node.Position {
	return n.position
}

func (n Mod) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Mod) Walk(v node.Visitor) {
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
