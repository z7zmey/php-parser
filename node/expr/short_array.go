package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortArray struct {
	attributes map[string]interface{}
	position   *node.Position
	items      []node.Node
}

func NewShortArray(items []node.Node) node.Node {
	return &ShortArray{
		map[string]interface{}{},
		nil,
		items,
	}
}

func (n ShortArray) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ShortArray) Position() *node.Position {
	return n.position
}

func (n ShortArray) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShortArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.items != nil {
		vv := v.GetChildrenVisitor("items")
		for _, nn := range n.items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
