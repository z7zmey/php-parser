package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortList struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	items      []node.Node
}

func NewShortList(items []node.Node) node.Node {
	return ShortList{
		"ShortList",
		map[string]interface{}{},
		nil,
		items,
	}
}

func (n ShortList) Name() string {
	return "ShortList"
}

func (n ShortList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ShortList) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ShortList) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n ShortList) Position() *node.Position {
	return n.position
}

func (n ShortList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShortList) Walk(v node.Visitor) {
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
