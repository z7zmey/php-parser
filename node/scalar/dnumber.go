package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Dnumber struct {
	position *node.Position
	comments *[]comment.Comment
	Value    string
}

func NewDnumber(Value string) node.Node {
	return &Dnumber{
		nil,
		nil,
		Value,
	}
}

func (n Dnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n Dnumber) Position() *node.Position {
	return n.position
}

func (n Dnumber) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Dnumber) Comments() *[]comment.Comment {
	return n.comments
}

func (n Dnumber) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Dnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
