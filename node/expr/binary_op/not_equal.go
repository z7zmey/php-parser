package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(variable node.Node, expression node.Node) node.Node {
	return &NotEqual{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n NotEqual) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NotEqual) Position() *node.Position {
	return n.position
}

func (n NotEqual) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n NotEqual) Walk(v node.Visitor) {
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
