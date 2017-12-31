package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewPrint(expression node.Node) node.Node {
	return Print{
		"Print",
		map[string]interface{}{},
		expression,
	}
}

func (n Print) Name() string {
	return "Print"
}

func (n Print) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Print) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Print) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Print) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
