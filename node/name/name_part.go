package name

import (
	"github.com/z7zmey/php-parser/node"
)

type NamePart struct {
	Value string
}

func NewNamePart(Value string) *NamePart {
	return &NamePart{
		Value,
	}
}

func (n *NamePart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
