package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Exit struct {
	Expr  node.Node
	IsDie bool
}

func NewExit(Expr node.Node, IsDie bool) *Exit {
	return &Exit{
		Expr,
		IsDie,
	}
}

func (n *Exit) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"IsDie": n.IsDie,
	}
}

func (n *Exit) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
