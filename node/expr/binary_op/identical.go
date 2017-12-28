package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Identical) Name() string {
	return "Identical"
}

type Identical struct {
	BinaryOp
}

func NewIdentical(variable node.Node, expression node.Node) node.Node {
	return Identical{
		BinaryOp{
			"BinaryIdentical",
			variable,
			expression,
		},
	}
}

func (n Identical) Walk(v node.Visitor) {
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
