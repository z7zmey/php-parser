package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// InlineHtml node
type InlineHtml struct {
	Comments []*comment.Comment
	Position *position.Position
	Value    string
}

// NewInlineHtml node constructor
func NewInlineHtml(Value string) *InlineHtml {
	return &InlineHtml{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *InlineHtml) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *InlineHtml) GetPosition() *position.Position {
	return n.Position
}

func (n *InlineHtml) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *InlineHtml) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *InlineHtml) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InlineHtml) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
