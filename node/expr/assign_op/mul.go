package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	AssignOp
}

func NewMul(Variable node.Node, Expression node.Node) node.Node {
	return &Mul{
		AssignOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Mul) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Mul) Position() *node.Position {
	return n.position
}

func (n Mul) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Mul) Walk(v node.Visitor) {
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
