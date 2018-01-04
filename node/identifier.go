package node

import (
	"github.com/z7zmey/php-parser/token"
)

type Identifier struct {
	attributes map[string]interface{}
	position   *Position
}

func NewIdentifier(token token.Token) Node {
	return &Identifier{
		map[string]interface{}{
			"Value": token.Value,
		},
		nil,
	}
}

func (n Identifier) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Identifier) Attribute(Key string) interface{} {
	return n.attributes[Key]
}

func (n Identifier) SetAttribute(Key string, Value interface{}) Node {
	n.attributes[Key] = Value
	return n
}

func (n Identifier) Position() *Position {
	return n.position
}

func (n Identifier) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
