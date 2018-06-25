package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Dnumber node
type Dnumber struct {
	Comments []*comment.Comment
	Position *position.Position
	Value    string
}

// NewDnumber node constructor
func NewDnumber(Value string) *Dnumber {
	return &Dnumber{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *Dnumber) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Dnumber) GetPosition() *position.Position {
	return n.Position
}

func (n *Dnumber) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Dnumber) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Dnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Dnumber) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
