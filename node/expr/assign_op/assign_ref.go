package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignRef struct {
	AssignOp
}

func NewAssignRef(Variable node.Node, Expression node.Node) *AssignRef {
	return &AssignRef{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *AssignRef) Attributes() map[string]interface{} {
	return nil
}

func (n *AssignRef) Walk(v node.Visitor) {
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
