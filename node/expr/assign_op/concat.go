package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	AssignOp
}

func NewConcat(Variable node.Node, Expression node.Node) *Concat {
	return &Concat{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *Concat) Attributes() map[string]interface{} {
	return nil
}

func (n *Concat) Walk(v node.Visitor) {
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
