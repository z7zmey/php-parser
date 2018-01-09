package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	AssignOp
}

func NewPlus(Variable node.Node, Expression node.Node) *Plus {
	return &Plus{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *Plus) Attributes() map[string]interface{} {
	return nil
}

func (n *Plus) Walk(v node.Visitor) {
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
