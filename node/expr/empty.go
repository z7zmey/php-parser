package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Empty struct {
	Expr node.Node
}

func NewEmpty(Expression node.Node) *Empty {
	return &Empty{
		Expression,
	}
}

func (n *Empty) Attributes() map[string]interface{} {
	return nil
}

func (n *Empty) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
