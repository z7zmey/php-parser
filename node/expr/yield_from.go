package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type YieldFrom struct {
	Expr node.Node
}

func NewYieldFrom(Expression node.Node) *YieldFrom {
	return &YieldFrom{
		Expression,
	}
}

func (n *YieldFrom) Attributes() map[string]interface{} {
	return nil
}

func (n *YieldFrom) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
