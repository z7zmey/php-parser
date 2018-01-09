package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	AssignOp
}

func NewMod(Variable node.Node, Expression node.Node) *Mod {
	return &Mod{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *Mod) Attributes() map[string]interface{} {
	return nil
}

func (n *Mod) Walk(v node.Visitor) {
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
