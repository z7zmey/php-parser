package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n SmallerOrEqual) Name() string {
	return "SmallerOrEqual"
}

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(variable node.Node, expression node.Node) node.Node {
	return SmallerOrEqual{
		BinaryOp{
			"BinarySmallerOrEqual",
			variable,
			expression,
		},
	}
}

func (n SmallerOrEqual) Walk(v node.Visitor) {
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
