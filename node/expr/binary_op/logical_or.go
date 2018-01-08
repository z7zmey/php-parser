package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalOr struct {
	BinaryOp
}

func NewLogicalOr(Variable node.Node, Expression node.Node) *LogicalOr {
	return &LogicalOr{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *LogicalOr) Attributes() map[string]interface{} {
	return nil
}

func (n *LogicalOr) Position() *node.Position {
	return n.position
}

func (n *LogicalOr) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *LogicalOr) Walk(v node.Visitor) {
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
