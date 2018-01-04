package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassConstList struct {
	attributes map[string]interface{}
	position   *node.Position
	modifiers  []node.Node
	consts     []node.Node
}

func NewClassConstList(modifiers []node.Node, consts []node.Node) node.Node {
	return &ClassConstList{
		map[string]interface{}{},
		nil,
		modifiers,
		consts,
	}
}

func (n ClassConstList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClassConstList) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ClassConstList) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ClassConstList) Position() *node.Position {
	return n.position
}

func (n ClassConstList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClassConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.consts != nil {
		vv := v.GetChildrenVisitor("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
