package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayItem struct {
	attributes map[string]interface{}
	position   *node.Position
	Key        node.Node
	Val        node.Node
}

func NewArrayItem(Key node.Node, Val node.Node, byRef bool) node.Node {
	return &ArrayItem{
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		Key,
		Val,
	}
}

func (n ArrayItem) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Val != nil {
		vv := v.GetChildrenVisitor("Val")
		n.Val.Walk(vv)
	}

	v.LeaveNode(n)
}
