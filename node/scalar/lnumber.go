package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Lnumber struct {
	position *node.Position
	comments *[]comment.Comment
	Value    string
}

func NewLnumber(Value string) node.Node {
	return &Lnumber{
		nil,
		nil,
		Value,
	}
}

func (n Lnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n Lnumber) Position() *node.Position {
	return n.position
}

func (n Lnumber) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Lnumber) Comments() *[]comment.Comment {
	return n.comments
}

func (n Lnumber) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
