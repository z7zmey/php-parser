package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ShortList struct {
	position *node.Position
	comments *[]comment.Comment
	Items    []node.Node
}

func NewShortList(Items []node.Node) *ShortList {
	return &ShortList{
		nil,
		nil,
		Items,
	}
}

func (n ShortList) Attributes() map[string]interface{} {
	return nil
}

func (n ShortList) Position() *node.Position {
	return n.position
}

func (n ShortList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShortList) Comments() *[]comment.Comment {
	return n.comments
}

func (n ShortList) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n ShortList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		vv := v.GetChildrenVisitor("Items")
		for _, nn := range n.Items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
