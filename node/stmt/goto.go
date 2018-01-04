package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Goto struct {
	attributes map[string]interface{}
	position   *node.Position
	label      node.Node
}

func NewGoto(label node.Node) node.Node {
	return Goto{
		map[string]interface{}{},
		nil,
		label,
	}
}

func (n Goto) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Goto) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Goto) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Goto) Position() *node.Position {
	return n.position
}

func (n Goto) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Goto) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.label != nil {
		vv := v.GetChildrenVisitor("label")
		n.label.Walk(vv)
	}

	v.LeaveNode(n)
}
