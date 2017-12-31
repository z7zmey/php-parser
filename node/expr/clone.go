package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Clone struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewClone(expression node.Node) node.Node {
	return Clone{
		"Clone",
		map[string]interface{}{},
		expression,
	}
}

func (n Clone) Name() string {
	return "Clone"
}

func (n Clone) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Clone) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Clone) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Clone) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
