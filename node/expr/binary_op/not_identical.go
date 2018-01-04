package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotIdentical struct {
	BinaryOp
}

func NewNotIdentical(variable node.Node, expression node.Node) node.Node {
	return NotIdentical{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n NotIdentical) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NotIdentical) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n NotIdentical) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n NotIdentical) Position() *node.Position {
	return n.position
}

func (n NotIdentical) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n NotIdentical) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
