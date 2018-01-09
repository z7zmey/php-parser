package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	BinaryOp
}

func NewBitwiseAnd(Variable node.Node, Expression node.Node) *BitwiseAnd {
	return &BitwiseAnd{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *BitwiseAnd) Attributes() map[string]interface{} {
	return nil
}

func (n *BitwiseAnd) Walk(v node.Visitor) {
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
