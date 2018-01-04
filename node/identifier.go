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
			"value": token.Value,
		},
		nil,
	}
}

func (n Identifier) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Identifier) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Identifier) SetAttribute(key string, value interface{}) Node {
	n.attributes[key] = value
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
