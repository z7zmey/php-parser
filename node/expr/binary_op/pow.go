package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	BinaryOp
}

func NewPow(Variable node.Node, Expression node.Node) *Pow {
	return &Pow{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *Pow) Attributes() map[string]interface{} {
	return nil
}

func (n *Pow) Walk(v node.Visitor) {
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
