package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalAnd struct {
	BinaryOp
}

func NewLogicalAnd(Variable node.Node, Expression node.Node) node.Node {
	return &LogicalAnd{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n LogicalAnd) Attributes() map[string]interface{} {
	return nil
}

func (n LogicalAnd) Position() *node.Position {
	return n.position
}

func (n LogicalAnd) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n LogicalAnd) Walk(v node.Visitor) {
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
