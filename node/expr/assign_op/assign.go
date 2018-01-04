package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Assign struct {
	AssignOp
}

func NewAssign(Variable node.Node, Expression node.Node) node.Node {
	return &Assign{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Assign) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Assign) Position() *node.Position {
	return n.position
}

func (n Assign) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Assign) Walk(v node.Visitor) {
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
