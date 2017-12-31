package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreDec struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	variable   node.Node
}

func NewPreDec(variable node.Node) node.Node {
	return PreDec{
		"PreDec",
		map[string]interface{}{},
		nil,
		variable,
	}
}

func (n PreDec) Name() string {
	return "PreDec"
}

func (n PreDec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PreDec) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PreDec) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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
