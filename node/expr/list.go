package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type List struct {
	attributes map[string]interface{}
	position   *node.Position
	Items      []node.Node
}

func NewList(Items []node.Node) node.Node {
	return &List{
		map[string]interface{}{},
		nil,
		Items,
	}
}

func (n List) Attributes() map[string]interface{} {
	return n.attributes
}

func (n List) Position() *node.Position {
	return n.position
}

func (n List) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n List) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		vv := v.GetChildrenVisitor("Items")
		for _, nn := range n.Items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
