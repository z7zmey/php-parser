package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type New struct {
	position  *node.Position
	comments  *[]comment.Comment
	Class     node.Node
	Arguments []node.Node
}

func NewNew(Class node.Node, Arguments []node.Node) node.Node {
	return &New{
		nil,
		nil,
		Class,
		Arguments,
	}
}

func (n New) Attributes() map[string]interface{} {
	return nil
}

func (n New) Position() *node.Position {
	return n.position
}

func (n New) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n New) Comments() *[]comment.Comment {
	return n.comments
}

func (n New) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n New) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
