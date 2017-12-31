package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassConstList struct {
	name       string
	attributes map[string]interface{}
	modifiers  []node.Node
	consts     []node.Node
}

func NewClassConstList(modifiers []node.Node, consts []node.Node) node.Node {
	return ClassConstList{
		"ClassConstList",
		map[string]interface{}{},
		modifiers,
		consts,
	}
}

func (n ClassConstList) Name() string {
	return "ClassConstList"
}

func (n ClassConstList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClassConstList) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ClassConstList) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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
