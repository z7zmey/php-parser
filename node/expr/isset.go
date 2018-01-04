package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	attributes map[string]interface{}
	position   *node.Position
	Variables  []node.Node
}

func NewIsset(Variables []node.Node) node.Node {
	return &Isset{
		map[string]interface{}{},
		nil,
		Variables,
	}
}

func (n Isset) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Isset) Position() *node.Position {
	return n.position
}

func (n Isset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Isset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variables != nil {
		vv := v.GetChildrenVisitor("Variables")
		for _, nn := range n.Variables {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
