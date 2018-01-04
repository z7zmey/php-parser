package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Break struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}

func NewBreak(expr node.Node) node.Node {
	return &Break{
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Break) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Break) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Break) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Break) Position() *node.Position {
	return n.position
}

func (n Break) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Break) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
