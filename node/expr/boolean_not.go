package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanNot struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewBooleanNot(expression node.Node) node.Node {
	return BooleanNot{
		"BooleanNot",
		map[string]interface{}{},
		nil,
		expression,
	}
}

func (n BooleanNot) Name() string {
	return "BooleanNot"
}

func (n BooleanNot) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BooleanNot) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BooleanNot) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n BooleanNot) Position() *node.Position {
	return n.position
}

func (n BooleanNot) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BooleanNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
