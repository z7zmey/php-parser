package node

import (
	"github.com/z7zmey/php-parser/token"
)

type Identifier struct {
	name       string
	attributes map[string]interface{}
}

func NewIdentifier(token token.Token) Node {
	return Identifier{
		"Identifier",
		map[string]interface{}{
			"value": token.Value,
		},
	}
}

func (n Identifier) Name() string {
	return "Identifier"
}

func (n Identifier) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Identifier) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Identifier) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
