package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ClusureUse struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
}

func NewClusureUse(variable node.Node, byRef bool) node.Node {
	return ClusureUse{
		"ClusureUse",
		map[string]interface{}{
			"byRef": byRef,
		},
		nil,
		variable,
	}
}

func (n ClusureUse) Name() string {
	return "ClusureUse"
}

func (n ClusureUse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClusureUse) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ClusureUse) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ClusureUse) Position() *node.Position {
	return n.position
}

func (n ClusureUse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClusureUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
