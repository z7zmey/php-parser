package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type IncludeOnce struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewIncludeOnce(expression node.Node) node.Node {
	return &IncludeOnce{
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n IncludeOnce) Attributes() map[string]interface{} {
	return n.attributes
}

func (n IncludeOnce) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n IncludeOnce) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n IncludeOnce) Position() *node.Position {
	return n.position
}

func (n IncludeOnce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n IncludeOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
