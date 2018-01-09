package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	Expr node.Node
}

func NewUnaryMinus(Expression node.Node) *UnaryMinus {
	return &UnaryMinus{
		Expression,
	}
}

func (n *UnaryMinus) Attributes() map[string]interface{} {
	return nil
}

func (n *UnaryMinus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
