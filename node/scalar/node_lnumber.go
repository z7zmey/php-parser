package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Lnumber node
type Lnumber struct {
	Comments []*comment.Comment
	Position *position.Position
	Value    string
}

// NewLnumber node constructor
func NewLnumber(Value string) *Lnumber {
	return &Lnumber{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *Lnumber) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Lnumber) GetPosition() *position.Position {
	return n.Position
}

func (n *Lnumber) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Lnumber) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Lnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Lnumber) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
