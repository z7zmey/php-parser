package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	Expr node.Node
}

func NewEval(Expression node.Node) *Eval {
	return &Eval{
		Expression,
	}
}

func (n *Eval) Attributes() map[string]interface{} {
	return nil
}

func (n *Eval) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
