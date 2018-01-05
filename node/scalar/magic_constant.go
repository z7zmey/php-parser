package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type MagicConstant struct {
	position *node.Position
	comments *[]comment.Comment
	Value    string
}

func NewMagicConstant(Value string) node.Node {
	return &MagicConstant{
		nil,
		nil,
		Value,
	}
}

func (n MagicConstant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n MagicConstant) Position() *node.Position {
	return n.position
}

func (n MagicConstant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n MagicConstant) Comments() *[]comment.Comment {
	return n.comments
}

func (n MagicConstant) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n MagicConstant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
