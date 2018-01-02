package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Unset struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	vars       []node.Node
}

func NewUnset(vars []node.Node) node.Node {
	return Unset{
		"Unset",
		map[string]interface{}{},
		nil,
		vars,
	}
}

func (n Unset) Name() string {
	return "Unset"
}

func (n Unset) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Unset) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Unset) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Unset) Position() *node.Position {
	return n.position
}

func (n Unset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Unset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.GetChildrenVisitor("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
