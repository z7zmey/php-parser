package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
}

func NewPreDec(variable node.Node) node.Node {
	return &PreDec{
		map[string]interface{}{},
		nil,
		variable,
	}
}

func (n PreDec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PreDec) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PreDec) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n PreDec) Position() *node.Position {
	return n.position
}

func (n PreDec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PreDec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
