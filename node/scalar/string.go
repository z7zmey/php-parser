package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type String struct {
	position *node.Position
	comments []comment.Comment
	Value    string
}

func NewString(Value string) *String {
	return &String{
		nil,
		nil,
		Value,
	}
}

func (n *String) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *String) Position() *node.Position {
	return n.position
}

func (n *String) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *String) Comments() []comment.Comment {
	return n.comments
}

func (n *String) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *String) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
