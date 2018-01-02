package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Ternary struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	condition  node.Node
	ifTrue     node.Node
	ifFalse    node.Node
}

func NewTernary(condition node.Node, ifTrue node.Node, ifFalse node.Node) node.Node {
	return Ternary{
		"Ternary",
		map[string]interface{}{},
		nil,
		condition,
		ifTrue,
		ifFalse,
	}
}

func (n Ternary) Name() string {
	return "Ternary"
}

func (n Ternary) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Ternary) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Ternary) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Ternary) Position() *node.Position {
	return n.position
}

func (n Ternary) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Ternary) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.condition != nil {
		vv := v.GetChildrenVisitor("condition")
		n.condition.Walk(vv)
	}

	if n.ifTrue != nil {
		vv := v.GetChildrenVisitor("ifTrue")
		n.ifTrue.Walk(vv)
	}

	if n.ifFalse != nil {
		vv := v.GetChildrenVisitor("ifFalse")
		n.ifFalse.Walk(vv)
	}

	v.LeaveNode(n)
}
