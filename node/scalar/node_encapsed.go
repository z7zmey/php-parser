package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Encapsed node
type Encapsed struct {
	Comments []*comment.Comment
	Position *position.Position
	Parts    []node.Node
}

// NewEncapsed node constructor
func NewEncapsed(Parts []node.Node) *Encapsed {
	return &Encapsed{
		Parts: Parts,
	}
}

// SetPosition sets node position
func (n *Encapsed) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Encapsed) GetPosition() *position.Position {
	return n.Position
}

func (n *Encapsed) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Encapsed) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Encapsed) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Encapsed) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		v.EnterChildList("Parts", n)
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Parts", n)
	}

	v.LeaveNode(n)
}
