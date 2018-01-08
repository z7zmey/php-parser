package name

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type NamePart struct {
	position *node.Position
	comments *[]comment.Comment
	Value    string
}

func NewNamePart(Value string) *NamePart {
	return &NamePart{
		nil,
		nil,
		Value,
	}
}

func (n NamePart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n NamePart) Position() *node.Position {
	return n.position
}

func (n NamePart) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n NamePart) Comments() *[]comment.Comment {
	return n.comments
}

func (n NamePart) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
