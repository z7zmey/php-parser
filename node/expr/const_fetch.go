package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ConstFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	constant   node.Node
}

func NewConstFetch(constant node.Node) node.Node {
	return ConstFetch{
		map[string]interface{}{},
		nil,
		constant,
	}
}

func (n ConstFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ConstFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ConstFetch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ConstFetch) Position() *node.Position {
	return n.position
}

func (n ConstFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ConstFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constant != nil {
		vv := v.GetChildrenVisitor("constant")
		n.constant.Walk(vv)
	}

	v.LeaveNode(n)
}
