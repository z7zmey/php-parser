package name

import (
	"github.com/z7zmey/php-parser/node"
)

type NamePart struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
}

func NewNamePart(value string) node.Node {
	return NamePart{
		"NamePart",
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n NamePart) Name() string {
	return "NamePart"
}

func (n NamePart) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NamePart) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n NamePart) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n NamePart) Position() *node.Position {
	return n.position
}

func (n NamePart) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
