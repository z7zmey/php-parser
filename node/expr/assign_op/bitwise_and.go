package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return &BitwiseAnd{
		AssignOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n BitwiseAnd) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseAnd) Position() *node.Position {
	return n.position
}

func (n BitwiseAnd) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseAnd) Walk(v node.Visitor) {
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
