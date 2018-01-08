package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type EncapsedStringPart struct {
	position *node.Position
	Value    string
}

func NewEncapsedStringPart(Value string) *EncapsedStringPart {
	return &EncapsedStringPart{
		nil,
		Value,
	}
}

func (n *EncapsedStringPart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *EncapsedStringPart) Position() *node.Position {
	return n.position
}

func (n *EncapsedStringPart) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *EncapsedStringPart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
