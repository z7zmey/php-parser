package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ShortArray struct {
	position *node.Position
	comments []comment.Comment
	Items    []node.Node
}

func NewShortArray(Items []node.Node) *ShortArray {
	return &ShortArray{
		nil,
		nil,
		Items,
	}
}

func (n *ShortArray) Attributes() map[string]interface{} {
	return nil
}

func (n *ShortArray) Position() *node.Position {
	return n.position
}

func (n *ShortArray) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ShortArray) Comments() []comment.Comment {
	return n.comments
}

func (n *ShortArray) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ShortArray) Walk(v node.Visitor) {
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
