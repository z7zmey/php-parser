package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Minus struct {
	BinaryOp
}

func NewMinus(Variable node.Node, Expression node.Node) node.Node {
	return &Minus{
		BinaryOp{
			map[string]interface{}{},
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Minus) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Minus) Position() *node.Position {
	return n.position
}

func (n Minus) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Minus) Walk(v node.Visitor) {
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
