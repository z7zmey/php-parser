package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(Variable node.Node, Expression node.Node) node.Node {
	return &Coalesce{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Coalesce) Attributes() map[string]interface{} {
	return nil
}

func (n Coalesce) Position() *node.Position {
	return n.position
}

func (n Coalesce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Coalesce) Walk(v node.Visitor) {
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
