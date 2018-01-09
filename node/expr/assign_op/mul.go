package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	AssignOp
}

func NewMul(Variable node.Node, Expression node.Node) *Mul {
	return &Mul{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *Mul) Attributes() map[string]interface{} {
	return nil
}

func (n *Mul) Walk(v node.Visitor) {
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
