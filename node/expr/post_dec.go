package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostDec struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	variable   node.Node
}

func NewPostDec(variable node.Node) node.Node {
	return PostDec{
		"PostDec",
		map[string]interface{}{},
		nil,
		variable,
	}
}

func (n PostDec) Name() string {
	return "PostDec"
}

func (n PostDec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PostDec) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PostDec) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n PostDec) Position() *node.Position {
	return n.position
}

func (n PostDec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PostDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
