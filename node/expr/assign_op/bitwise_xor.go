package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	AssignOp
}

func NewBitwiseXor(Variable node.Node, Expression node.Node) node.Node {
	return &BitwiseXor{
		AssignOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n BitwiseXor) Attributes() map[string]interface{} {
	return nil
}

func (n BitwiseXor) Position() *node.Position {
	return n.position
}

func (n BitwiseXor) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseXor) Walk(v node.Visitor) {
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
