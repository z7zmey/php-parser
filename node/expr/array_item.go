package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayItem struct {
	position *node.Position
	ByRef    bool
	Key      node.Node
	Val      node.Node
}

func NewArrayItem(Key node.Node, Val node.Node, ByRef bool) node.Node {
	return &ArrayItem{
		nil,
		ByRef,
		Key,
		Val,
	}
}

func (n ArrayItem) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
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
