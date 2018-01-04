package name

import (
	"github.com/z7zmey/php-parser/node"
)

type Name struct {
	attributes map[string]interface{}
	position   *node.Position
	parts      []node.Node
}

func NewName(parts []node.Node) node.Node {
	return &Name{
		map[string]interface{}{},
		nil,
		parts,
	}
}

func (n Name) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Name) Position() *node.Position {
	return n.position
}

func (n Name) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Name) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.parts != nil {
		vv := v.GetChildrenVisitor("parts")
		for _, nn := range n.parts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
