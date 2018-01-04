package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShortList struct {
	attributes map[string]interface{}
	position   *node.Position
	Items      []node.Node
}

func NewShortList(Items []node.Node) node.Node {
	return &ShortList{
		map[string]interface{}{},
		nil,
		Items,
	}
}

func (n ShortList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ShortList) Position() *node.Position {
	return n.position
}

func (n ShortList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShortList) Walk(v node.Visitor) {
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
