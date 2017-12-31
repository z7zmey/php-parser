package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostInc struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
}

func NewPostInc(variable node.Node) node.Node {
	return PostInc{
		"PostInc",
		map[string]interface{}{},
		variable,
	}
}

func (n PostInc) Name() string {
	return "PostInc"
}

func (n PostInc) Attributes() map[string]interface{} {
	return n.attributes
}

func (n PostInc) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n PostInc) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n PostInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
