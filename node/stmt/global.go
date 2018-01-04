package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	attributes map[string]interface{}
	position   *node.Position
	vars       []node.Node
}

func NewGlobal(vars []node.Node) node.Node {
	return &Global{
		map[string]interface{}{},
		nil,
		vars,
	}
}

func (n Global) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Global) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Global) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Global) Position() *node.Position {
	return n.position
}

func (n Global) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Global) Walk(v node.Visitor) {
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
