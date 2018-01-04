package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	BinaryOp
}

func NewConcat(variable node.Node, expression node.Node) node.Node {
	return Concat{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Concat) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Concat) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Concat) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Concat) Position() *node.Position {
	return n.position
}

func (n Concat) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Concat) Walk(v node.Visitor) {
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
