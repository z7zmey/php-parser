package node

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Identifier node
type Identifier struct {
	Comments []*comment.Comment
	Position *position.Position
	Value    string
}

// NewIdentifier node constructor
func NewIdentifier(Value string) *Identifier {
	return &Identifier{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *Identifier) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Identifier) GetPosition() *position.Position {
	return n.Position
}

func (n *Identifier) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Identifier) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Identifier) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Identifier) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
