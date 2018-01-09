package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(Variable node.Node, Expression node.Node) *BitwiseAnd {
	return &BitwiseAnd{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *BitwiseAnd) Attributes() map[string]interface{} {
	return nil
}

func (n *BitwiseAnd) Walk(v node.Visitor) {
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
