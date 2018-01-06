package assign_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	AssignOp
}

func NewConcat(Variable node.Node, Expression node.Node) node.Node {
	return &Concat{
		AssignOp{
			nil,
			nil,
			Variable,
			Expression,
		},
	}
}

func (n Concat) Attributes() map[string]interface{} {
	return nil
}

func (n Concat) Position() *node.Position {
	return n.position
}

func (n Concat) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Concat) Comments() *[]comment.Comment {
	return n.comments
}

func (n Concat) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Concat) Walk(v node.Visitor) {
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
