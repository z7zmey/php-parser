package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type New struct {
	attributes map[string]interface{}
	position   *node.Position
	class      node.Node
	arguments  []node.Node
}

func NewNew(class node.Node, arguments []node.Node) node.Node {
	return &New{
		map[string]interface{}{},
		nil,
		class,
		arguments,
	}
}

func (n New) Attributes() map[string]interface{} {
	return n.attributes
}

func (n New) Position() *node.Position {
	return n.position
}

func (n New) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n New) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.GetChildrenVisitor("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
