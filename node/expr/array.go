package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Array struct {
	attributes map[string]interface{}
	position   *node.Position
	items      []node.Node
}

func NewArray(items []node.Node) node.Node {
	return Array{
		map[string]interface{}{},
		nil,
		items,
	}
}

func (n Array) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Array) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Array) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Array) Position() *node.Position {
	return n.position
}

func (n Array) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
