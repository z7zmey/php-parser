package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Identical struct {
	BinaryOp
}

func NewIdentical(variable node.Node, expression node.Node) node.Node {
	return Identical{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Identical) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Identical) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Identical) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Identical) Position() *node.Position {
	return n.position
}

func (n Identical) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Identical) Walk(v node.Visitor) {
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
