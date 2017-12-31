package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	variables  []node.Node
}

func NewIsset(variables []node.Node) node.Node {
	return Isset{
		"Isset",
		map[string]interface{}{},
		nil,
		variables,
	}
}

func (n Isset) Name() string {
	return "Isset"
}

func (n Isset) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Isset) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Isset) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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

	if n.variables != nil {
		vv := v.GetChildrenVisitor("variables")
		for _, nn := range n.variables {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
