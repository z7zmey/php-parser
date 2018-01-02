package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Nop struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
}

func NewNop() node.Node {
	return Nop{
		"Nop",
		map[string]interface{}{},
		nil,
	}
}

func (n Nop) Name() string {
	return "Nop"
}

func (n Nop) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Nop) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Nop) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Nop) Position() *node.Position {
	return n.position
}

func (n Nop) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
