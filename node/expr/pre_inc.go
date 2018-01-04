package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
}

func NewPreInc(variable node.Node) node.Node {
	return PreInc{
		map[string]interface{}{},
		nil,
		variable,
	}
}

func (n PreInc) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PreInc) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PreInc) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n PreInc) Position() *node.Position {
	return n.position
}

func (n PreInc) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n PreInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
