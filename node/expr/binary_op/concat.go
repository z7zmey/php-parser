package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Concat) Name() string {
	return "Concat"
}

type Concat struct {
	BinaryOp
}

func NewConcat(variable node.Node, expression node.Node) node.Node {
	return Concat{
		BinaryOp{
			"BinaryConcat",
			variable,
			expression,
		},
	}
}

func (n Concat) Walk(v node.Visitor) {
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
