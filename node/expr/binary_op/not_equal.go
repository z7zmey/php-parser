package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(Variable node.Node, Expression node.Node) *NotEqual {
	return &NotEqual{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *NotEqual) Attributes() map[string]interface{} {
	return nil
}

func (n *NotEqual) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
