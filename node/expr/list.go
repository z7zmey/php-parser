package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type List struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	items      []node.Node
}

func NewList(items []node.Node) node.Node {
	return List{
		"List",
		map[string]interface{}{},
		nil,
		items,
	}
}

func (n List) Name() string {
	return "List"
}

func (n List) Attributes() map[string]interface{} {
	return n.attributes
}

func (n List) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n List) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n List) Position() *node.Position {
	return n.position
}

func (n List) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n List) Walk(v node.Visitor) {
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
