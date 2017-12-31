package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Continue struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewContinue(expr node.Node) node.Node {
	return Continue{
		"Continue",
		map[string]interface{}{},
		expr,
	}
}

func (n Continue) Name() string {
	return "Continue"
}

func (n Continue) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Continue) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Continue) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Continue) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
