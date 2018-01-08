package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type MagicConstant struct {
	position *node.Position
	Value    string
}

func NewMagicConstant(Value string) *MagicConstant {
	return &MagicConstant{
		nil,
		Value,
	}
}

func (n *MagicConstant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *MagicConstant) Position() *node.Position {
	return n.position
}

func (n *MagicConstant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *MagicConstant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
