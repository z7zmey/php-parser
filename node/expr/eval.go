package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewEval(expression node.Node) node.Node {
	return Eval{
		"Eval",
		map[string]interface{}{},
		expression,
	}
}

func (n Eval) Name() string {
	return "Eval"
}

func (n Eval) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Eval) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Eval) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Eval) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
