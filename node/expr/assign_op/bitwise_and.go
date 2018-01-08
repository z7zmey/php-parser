package assign_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type BitwiseAnd struct {
	AssignOp
}

func NewBitwiseAnd(Variable node.Node, Expression node.Node) *BitwiseAnd {
	return &BitwiseAnd{
		AssignOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n *BitwiseAnd) Attributes() map[string]interface{} {
	return nil
}

func (n *BitwiseAnd) Position() *node.Position {
	return n.position
}

func (n *BitwiseAnd) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *BitwiseAnd) Comments() *[]comment.Comment {
	return n.comments
}

func (n *BitwiseAnd) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *BitwiseAnd) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expression != nil {
		vv := v.GetChildrenVisitor("Expression")
		n.Expression.Walk(vv)
	}

	v.LeaveNode(n)
}
