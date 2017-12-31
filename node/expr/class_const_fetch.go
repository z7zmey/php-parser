package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ClassConstFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClassConstFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ClassConstFetch) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n ClassConstFetch) Position() *node.Position {
	return n.position
}

func (n ClassConstFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClassConstFetch) Name() string {
	return "ClassConstFetch"
}

type ClassConstFetch struct {
	name         string
	attributes   map[string]interface{}
	position *node.Position
	class        node.Node
	constantName node.Node
}

func NewClassConstFetch(class node.Node, constantName node.Node) node.Node {
	return ClassConstFetch{
		"ClassConstFetch",
		map[string]interface{}{},
		nil,
		class,
		constantName,
	}
}

func (n ClassConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constantName != nil {
		vv := v.GetChildrenVisitor("constantName")
		n.constantName.Walk(vv)
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	v.LeaveNode(n)
}
