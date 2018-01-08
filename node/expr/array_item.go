package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ArrayItem struct {
	position *node.Position
	comments []comment.Comment
	ByRef    bool
	Key      node.Node
	Val      node.Node
}

func NewArrayItem(Key node.Node, Val node.Node, ByRef bool) *ArrayItem {
	return &ArrayItem{
		nil,
		nil,
		ByRef,
		Key,
		Val,
	}
}

func (n *ArrayItem) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
}

func (n *ArrayItem) Position() *node.Position {
	return n.position
}

func (n *ArrayItem) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ArrayItem) Comments() []comment.Comment {
	return n.comments
}

func (n *ArrayItem) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ArrayItem) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Val != nil {
		vv := v.GetChildrenVisitor("Val")
		n.Val.Walk(vv)
	}

	v.LeaveNode(n)
}
