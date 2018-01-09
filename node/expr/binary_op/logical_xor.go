package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalXor struct {
	BinaryOp
}

func NewLogicalXor(Variable node.Node, Expression node.Node) *LogicalXor {
	return &LogicalXor{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *LogicalXor) Attributes() map[string]interface{} {
	return nil
}

func (n *LogicalXor) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
