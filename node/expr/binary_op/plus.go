package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	BinaryOp
}

func NewPlus(Variable node.Node, Expression node.Node) *Plus {
	return &Plus{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *Plus) Attributes() map[string]interface{} {
	return nil
}

func (n *Plus) Walk(v node.Visitor) {
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
