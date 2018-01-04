package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalXor struct {
	BinaryOp
}

func NewLogicalXor(Variable node.Node, Expression node.Node) node.Node {
	return &LogicalXor{
		BinaryOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n LogicalXor) Attributes() map[string]interface{} {
	return n.attributes
}

func (n LogicalXor) Position() *node.Position {
	return n.position
}

func (n LogicalXor) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n LogicalXor) Walk(v node.Visitor) {
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
