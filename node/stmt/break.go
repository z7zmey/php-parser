package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Break struct {
	Expr node.Node
}

func NewBreak(Expr node.Node) *Break {
	return &Break{
		Expr,
	}
}

func (n *Break) Attributes() map[string]interface{} {
	return nil
}

func (n *Break) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
