package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Property struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
	expr       node.Node
}

func NewProperty(variable node.Node, expr node.Node, phpDocComment string) node.Node {
	return Property{
		"Property",
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
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

func (n Property) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Property) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Property) Position() *node.Position {
	return n.position
}

func (n Property) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
