package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Equal) Name() string {
	return "Equal"
}

type Equal struct {
	BinaryOp
}

func NewEqual(variable node.Node, expression node.Node) node.Node {
	return Equal{
		BinaryOp{
			"BinaryEqual",
			variable,
			expression,
		},
	}
}

func (n Equal) Walk(v node.Visitor) {
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
