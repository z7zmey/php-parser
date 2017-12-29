package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	BinaryOp
}

func NewBitwiseOr(variable node.Node, expression node.Node) node.Node {
	return BitwiseOr{
		BinaryOp{
			"BinaryBitwiseOr",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n BitwiseOr) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseOr) Name() string {
	return "BitwiseOr"
}

func (n BitwiseOr) Walk(v node.Visitor) {
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
