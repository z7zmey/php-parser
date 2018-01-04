package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type EncapsedStringPart struct {
	attributes map[string]interface{}
	position   *node.Position
}

func NewEncapsedStringPart(value string) node.Node {
	return &EncapsedStringPart{
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n EncapsedStringPart) Attributes() map[string]interface{} {
	return n.attributes
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
