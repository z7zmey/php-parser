package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Equal struct {
	BinaryOp
}

func NewEqual(Variable node.Node, Expression node.Node) node.Node {
	return &Equal{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Equal) Attributes() map[string]interface{} {
	return nil
}

func (n Equal) Position() *node.Position {
	return n.position
}

func (n Equal) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Equal) Walk(v node.Visitor) {
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
