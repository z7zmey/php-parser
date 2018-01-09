package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type MagicConstant struct {
	Value string
}

func NewMagicConstant(Value string) *MagicConstant {
	return &MagicConstant{
		Value,
	}
}

func (n *MagicConstant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *MagicConstant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
