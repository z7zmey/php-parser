package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(variable node.Node, expression node.Node) node.Node {
	return &SmallerOrEqual{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n SmallerOrEqual) Attributes() map[string]interface{} {
	return n.attributes
}

func (n SmallerOrEqual) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n SmallerOrEqual) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n SmallerOrEqual) Position() *node.Position {
	return n.position
}

func (n SmallerOrEqual) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n SmallerOrEqual) Walk(v node.Visitor) {
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
