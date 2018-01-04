package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	AssignOp
}

func NewMinus(Variable node.Node, Expression node.Node) node.Node {
	return &Minus{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Minus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Minus) Position() *node.Position {
	return n.position
}

func (n Minus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Minus) Walk(v node.Visitor) {
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
