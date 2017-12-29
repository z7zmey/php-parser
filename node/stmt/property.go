package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Property struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
	expr       node.Node
}

func NewProperty(variable node.Node, expr node.Node) node.Node {
	return Property{
		"Property",
		map[string]interface{}{},
		variable,
		expr,
	}
}
func (n Property) Name() string {
	return "Property"
}

func (n Property) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Property) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
