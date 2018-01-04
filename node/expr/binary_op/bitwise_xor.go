package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	BinaryOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseXor) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n BitwiseXor) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n BitwiseXor) Position() *node.Position {
	return n.position
}

func (n BitwiseXor) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseXor) Walk(v node.Visitor) {
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
