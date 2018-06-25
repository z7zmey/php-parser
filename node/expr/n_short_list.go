package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ShortList node
type ShortList struct {
	Comments []*comment.Comment
	Position *position.Position
	Items    []node.Node
}

// NewShortList node constructor
func NewShortList(Items []node.Node) *ShortList {
	return &ShortList{
		Items: Items,
	}
}

// SetPosition sets node position
func (n *ShortList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ShortList) GetPosition() *position.Position {
	return n.Position
}

func (n *ShortList) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *ShortList) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *ShortList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShortList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		v.EnterChildList("Items", n)
		for _, nn := range n.Items {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Items", n)
	}

	v.LeaveNode(n)
}
