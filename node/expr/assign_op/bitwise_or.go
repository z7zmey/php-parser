package assign_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type BitwiseOr struct {
	AssignOp
}

func NewBitwiseOr(Variable node.Node, Expression node.Node) node.Node {
	return &BitwiseOr{
		AssignOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n BitwiseOr) Attributes() map[string]interface{} {
	return nil
}

func (n BitwiseOr) Position() *node.Position {
	return n.position
}

func (n BitwiseOr) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n BitwiseOr) Comments() *[]comment.Comment {
	return n.comments
}

func (n BitwiseOr) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n BitwiseOr) Walk(v node.Visitor) {
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
