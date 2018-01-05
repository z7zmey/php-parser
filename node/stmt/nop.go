package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Nop struct {
	position *node.Position
	comments *[]comment.Comment
}

func NewNop() node.Node {
	return &Nop{
		nil,
		nil,
	}
}

func (n Nop) Attributes() map[string]interface{} {
	return nil
}

func (n Nop) Position() *node.Position {
	return n.position
}

func (n Nop) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Nop) Comments() *[]comment.Comment {
	return n.comments
}

func (n Nop) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
