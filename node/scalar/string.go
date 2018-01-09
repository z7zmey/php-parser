package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type String struct {
	Value string
}

func NewString(Value string) *String {
	return &String{
		Value,
	}
}

func (n *String) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *String) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
