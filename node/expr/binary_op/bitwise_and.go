package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BitwiseAnd) Name() string {
	return "BitwiseAnd"
}

type BitwiseAnd struct {
	BinaryOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		BinaryOp{
			"BinaryBitwiseAnd",
			variable,
			expression,
		},
	}
}

func (n BitwiseAnd) Walk(v node.Visitor) {
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
