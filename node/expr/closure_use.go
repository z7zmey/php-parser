package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ClusureUse struct {
	position *node.Position
	comments *[]comment.Comment
	ByRef    bool
	Variable node.Node
}

func NewClusureUse(Variable node.Node, ByRef bool) *ClusureUse {
	return &ClusureUse{
		nil,
		nil,
		ByRef,
		Variable,
	}
}

func (n *ClusureUse) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef": n.ByRef,
	}
}

func (n *ClusureUse) Position() *node.Position {
	return n.position
}

func (n *ClusureUse) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ClusureUse) Comments() *[]comment.Comment {
	return n.comments
}

func (n *ClusureUse) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *ClusureUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
