package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Array struct {
	name       string
	attributes map[string]interface{}
	items      []node.Node
}

func NewArray(items []node.Node) node.Node {
	return Array{
		"Array",
		map[string]interface{}{},
		items,
	}
}

func (n Array) Name() string {
	return "Array"
}

func (n Array) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Array) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Array) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Array) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.items != nil {
		vv := v.GetChildrenVisitor("items")
		for _, nn := range n.items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
