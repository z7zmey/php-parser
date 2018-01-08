package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Smaller struct {
	BinaryOp
}

func NewSmaller(Variable node.Node, Expression node.Node) *Smaller {
	return &Smaller{
		BinaryOp{
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *Smaller) Attributes() map[string]interface{} {
	return nil
}

func (n *Smaller) Position() *node.Position {
	return n.position
}

func (n *Smaller) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Smaller) Walk(v node.Visitor) {
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
