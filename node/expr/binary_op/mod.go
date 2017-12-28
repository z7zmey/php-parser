package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Mod) Name() string {
	return "Mod"
}

type Mod struct {
	BinaryOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		BinaryOp{
			"BinaryMod",
			variable,
			expression,
		},
	}
}

func (n Mod) Walk(v node.Visitor) {
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
