package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Identical struct {
	BinaryOp
}

func NewIdentical(variable node.Node, expression node.Node) node.Node {
	return Identical{
		BinaryOp{
			"BinaryIdentical",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Identical) Name() string {
	return "Identical"
}

func (n Identical) Attributes() map[string]interface{} {
	return n.attributes
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
