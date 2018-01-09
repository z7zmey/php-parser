package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Div struct {
	BinaryOp
}

func NewDiv(Variable node.Node, Expression node.Node) *Div {
	return &Div{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *Div) Attributes() map[string]interface{} {
	return nil
}

func (n *Div) Walk(v node.Visitor) {
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
