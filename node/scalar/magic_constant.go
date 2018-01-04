package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type MagicConstant struct {
	attributes map[string]interface{}
	position   *node.Position
}

func NewMagicConstant(Value string) node.Node {
	return &MagicConstant{
		map[string]interface{}{
			"Value": Value,
		},
		nil,
	}
}

func (n MagicConstant) Attributes() map[string]interface{} {
	return n.attributes
}

func (n MagicConstant) Position() *node.Position {
	return n.position
}

func (n MagicConstant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n MagicConstant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
