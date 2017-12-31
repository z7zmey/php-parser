package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PreInc struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
}

func NewPreInc(variable node.Node) node.Node {
	return PreInc{
		"PreInc",
		map[string]interface{}{},
		variable,
	}
}

func (n PreInc) Name() string {
	return "PreInc"
}

func (n PreInc) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PreInc) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PreInc) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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
