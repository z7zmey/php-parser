package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	BinaryOp
}

func NewMod(Variable node.Node, Expression node.Node) *Mod {
	return &Mod{
		BinaryOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *Mod) Attributes() map[string]interface{} {
	return nil
}

func (n *Mod) Position() *node.Position {
	return n.position
}

func (n *Mod) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Mod) Comments() []comment.Comment {
	return n.comments
}

func (n *Mod) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Mod) Walk(v node.Visitor) {
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
