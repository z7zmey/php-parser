package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type RequireOnce struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewRequireOnce(expression node.Node) node.Node {
	return RequireOnce{
		"RequireOnce",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n RequireOnce) Name() string {
	return "RequireOnce"
}

func (n RequireOnce) Attributes() map[string]interface{} {
	return n.attributes
}

func (n RequireOnce) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n RequireOnce) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n RequireOnce) Position() *node.Position {
	return n.position
}

func (n RequireOnce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n RequireOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
