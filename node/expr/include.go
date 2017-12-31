package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Include struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewInclude(expression node.Node) node.Node {
	return Include{
		"Include",
		map[string]interface{}{},
		expression,
	}
}

func (n Include) Name() string {
	return "Include"
}

func (n Include) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Include) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Include) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Include) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
