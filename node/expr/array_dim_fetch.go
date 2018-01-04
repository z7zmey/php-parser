package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayDimFetch struct {
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
	dim        node.Node
}

func NewArrayDimFetch(variable node.Node, dim node.Node) node.Node {
	return &ArrayDimFetch{
		map[string]interface{}{},
		nil,
		variable,
		dim,
	}
}

func (n ArrayDimFetch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ArrayDimFetch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ArrayDimFetch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ArrayDimFetch) Position() *node.Position {
	return n.position
}

func (n ArrayDimFetch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ArrayDimFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.dim != nil {
		vv := v.GetChildrenVisitor("dim")
		n.dim.Walk(vv)
	}

	v.LeaveNode(n)
}
