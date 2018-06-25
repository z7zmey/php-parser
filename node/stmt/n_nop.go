package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Nop node
type Nop struct {
	Comments []*comment.Comment
	Position *position.Position
}

// NewNop node constructor
func NewNop() *Nop {
	return &Nop{}
}

// SetPosition sets node position
func (n *Nop) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Nop) GetPosition() *position.Position {
	return n.Position
}

func (n *Nop) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Nop) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Nop) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Nop) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
