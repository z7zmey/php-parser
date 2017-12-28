package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n NotEqual) Name() string {
	return "NotEqual"
}

type NotEqual struct {
	BinaryOp
}

func NewNotEqual(variable node.Node, expression node.Node) node.Node {
	return NotEqual{
		BinaryOp{
			"BinaryNotEqual",
			variable,
			expression,
		},
	}
}

func (n NotEqual) Walk(v node.Visitor) {
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
