package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Assign struct {
	AssignOp
}

func NewAssign(Variable node.Node, Expression node.Node) *Assign {
	return &Assign{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *Assign) Attributes() map[string]interface{} {
	return nil
}

func (n *Assign) Walk(v node.Visitor) {
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
