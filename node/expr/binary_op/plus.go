package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Plus struct {
	BinaryOp
}

func NewPlus(variable node.Node, expression node.Node) node.Node {
	return Plus{
		BinaryOp{
			"BinaryPlus",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Plus) Name() string {
	return "Plus"
}

func (n Plus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Plus) Walk(v node.Visitor) {
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
