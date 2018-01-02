package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayItem struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	key        node.Node
	val        node.Node
}

func NewArrayItem(key node.Node, val node.Node, byRef bool) node.Node {
	return ArrayItem{
		"ArrayItem",
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		key,
		val,
	}
}

func (n ArrayItem) Name() string {
	return "ArrayItem"
}

func (n ArrayItem) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ArrayItem) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ArrayItem) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ArrayItem) Position() *node.Position {
	return n.position
}

func (n ArrayItem) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ArrayItem) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.key != nil {
		vv := v.GetChildrenVisitor("key")
		n.key.Walk(vv)
	}

	if n.val != nil {
		vv := v.GetChildrenVisitor("val")
		n.val.Walk(vv)
	}

	v.LeaveNode(n)
}
