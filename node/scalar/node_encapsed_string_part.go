package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// EncapsedStringPart node
type EncapsedStringPart struct {
	Comments []*comment.Comment
	Position *position.Position
	Value    string
}

// NewEncapsedStringPart node constructor
func NewEncapsedStringPart(Value string) *EncapsedStringPart {
	return &EncapsedStringPart{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *EncapsedStringPart) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *EncapsedStringPart) GetPosition() *position.Position {
	return n.Position
}

func (n *EncapsedStringPart) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *EncapsedStringPart) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *EncapsedStringPart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *EncapsedStringPart) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
