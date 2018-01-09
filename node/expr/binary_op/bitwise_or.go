package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	BinaryOp
}

func NewBitwiseOr(Variable node.Node, Expression node.Node) *BitwiseOr {
	return &BitwiseOr{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *BitwiseOr) Attributes() map[string]interface{} {
	return nil
}

func (n *BitwiseOr) Walk(v node.Visitor) {
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
