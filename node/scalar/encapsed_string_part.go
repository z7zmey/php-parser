package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type EncapsedStringPart struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
}

func NewEncapsedStringPart(value string) node.Node {
	return EncapsedStringPart{
		"EncapsedStringPart",
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n EncapsedStringPart) Name() string {
	return "EncapsedStringPart"
}

func (n EncapsedStringPart) Attributes() map[string]interface{} {
	return n.attributes
}

func (n EncapsedStringPart) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n EncapsedStringPart) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n EncapsedStringPart) Position() *node.Position {
	return n.position
}

func (n EncapsedStringPart) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n EncapsedStringPart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
