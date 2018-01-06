package node

import "github.com/z7zmey/php-parser/comment"

type Nullable struct {
	position *Position
	comments *[]comment.Comment
	Expr     Node
}

func NewNullable(Expression Node) Node {
	return &Nullable{
		nil,
		nil,
		Expression,
	}
}

func (n Nullable) Attributes() map[string]interface{} {
	return nil
}

func (n Nullable) Position() *Position {
	return n.position
}

func (n Nullable) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Nullable) Comments() *[]comment.Comment {
	return n.comments
}

func (n Nullable) SetComments(c *[]comment.Comment) Node {
	n.comments = c
	return n
}

func (n Nullable) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
