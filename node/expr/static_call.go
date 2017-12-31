package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticCall struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	class      node.Node
	call       node.Node
	arguments  []node.Node
}

func NewStaticCall(class node.Node, call node.Node, arguments []node.Node) node.Node {
	return StaticCall{
		"StaticCall",
		map[string]interface{}{},
		nil,
		class,
		call,
		arguments,
	}
}

func (n StaticCall) Name() string {
	return "StaticCall"
}

func (n StaticCall) Attributes() map[string]interface{} {
	return n.attributes
}

func (n StaticCall) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n StaticCall) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n StaticCall) Position() *node.Position {
	return n.position
}

func (n StaticCall) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n StaticCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	if n.call != nil {
		vv := v.GetChildrenVisitor("call")
		n.call.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.GetChildrenVisitor("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
