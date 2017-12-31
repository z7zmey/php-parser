package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type YieldFrom struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewYieldFrom(expression node.Node) node.Node {
	return YieldFrom{
		"YieldFrom",
		map[string]interface{}{},
		expression,
	}
}

func (n YieldFrom) Name() string {
	return "YieldFrom"
}

func (n YieldFrom) Attributes() map[string]interface{} {
	return n.attributes
}

func (n YieldFrom) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n YieldFrom) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n YieldFrom) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
