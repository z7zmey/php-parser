package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseXor struct {
	BinaryOp
}

func NewBitwiseXor(variable node.Node, expression node.Node) node.Node {
	return BitwiseXor{
		BinaryOp{
			"BinaryBitwiseXor",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n BitwiseXor) Attributes() map[string]interface{} {
	return n.attributes
}

func (n BitwiseXor) Name() string {
	return "BitwiseXor"
}

func (n BitwiseXor) Walk(v node.Visitor) {
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
