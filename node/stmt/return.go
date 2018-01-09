package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Return struct {
	Expr node.Node
}

func NewReturn(Expr node.Node) *Return {
	return &Return{
		Expr,
	}
}

func (n *Return) Attributes() map[string]interface{} {
	return nil
}

func (n *Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
