package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Array struct {
	attributes map[string]interface{}
	position   *node.Position
	Items      []node.Node
}

func NewArray(Items []node.Node) node.Node {
	return &Array{
		map[string]interface{}{},
		nil,
		Items,
	}
}

func (n Array) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Array) Position() *node.Position {
	return n.position
}

func (n Array) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Array) Walk(v node.Visitor) {
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
