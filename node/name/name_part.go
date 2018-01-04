package name

import (
	"github.com/z7zmey/php-parser/node"
)

type NamePart struct {
	attributes map[string]interface{}
	position   *node.Position
}

func NewNamePart(Value string) node.Node {
	return &NamePart{
		map[string]interface{}{
			"Value": Value,
		},
		nil,
	}
}

func (n NamePart) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NamePart) Position() *node.Position {
	return n.position
}

func (n NamePart) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
