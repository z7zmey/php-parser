package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	BinaryOp
}

func NewBitwiseAnd(variable node.Node, expression node.Node) node.Node {
	return BitwiseAnd{
		BinaryOp{
			"BinaryBitwiseAnd",
			variable,
			expression,
		},
	}
}

func (n BitwiseAnd) Name() string {
	return "BitwiseAnd"
}

func (n BitwiseAnd) Attributes() map[string]interface{} {
	return nil
}

func (n BitwiseAnd) Walk(v node.Visitor) {
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
