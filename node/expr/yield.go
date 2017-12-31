package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Yield struct {
	name       string
	attributes map[string]interface{}
	key        node.Node
	value      node.Node
}

func NewYield(key node.Node, value node.Node) node.Node {
	return Yield{
		"Yield",
		map[string]interface{}{},
		key,
		value,
	}
}

func (n Yield) Name() string {
	return "Yield"
}

func (n Yield) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Yield) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Yield) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Yield) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.key != nil {
		vv := v.GetChildrenVisitor("key")
		n.key.Walk(vv)
	}

	if n.value != nil {
		vv := v.GetChildrenVisitor("value")
		n.value.Walk(vv)
	}

	v.LeaveNode(n)
}
