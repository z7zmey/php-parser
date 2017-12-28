package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n UnaryPlus) Name() string {
	return "UnaryPlus"
}

type UnaryPlus struct {
	name string
	expr node.Node
}

func NewUnaryPlus(expression node.Node) node.Node {
	return UnaryPlus{
		"UnaryPlus",
		expression,
	}
}

func (n UnaryPlus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
