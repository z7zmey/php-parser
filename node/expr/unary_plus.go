package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewUnaryPlus(expression node.Node) node.Node {
	return UnaryPlus{
		"UnaryPlus",
		map[string]interface{}{},
		expression,
	}
}

func (n UnaryPlus) Name() string {
	return "UnaryPlus"
}

func (n UnaryPlus) Attributes() map[string]interface{} {
	return n.attributes
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
