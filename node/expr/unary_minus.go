package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryMinus struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewUnaryMinus(expression node.Node) node.Node {
	return UnaryMinus{
		"UnaryMinus",
		map[string]interface{}{},
		expression,
	}
}

func (n UnaryMinus) Name() string {
	return "UnaryMinus"
}

func (n UnaryMinus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n UnaryMinus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
