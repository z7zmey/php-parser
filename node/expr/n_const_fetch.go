package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ConstFetch node
type ConstFetch struct {
	Comments []*comment.Comment
	Position *position.Position
	Constant node.Node
}

// NewConstFetch node constructor
func NewConstFetch(Constant node.Node) *ConstFetch {
	return &ConstFetch{
		Constant: Constant,
	}
}

// SetPosition sets node position
func (n *ConstFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ConstFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *ConstFetch) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *ConstFetch) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *ConstFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ConstFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Constant != nil {
		v.EnterChildNode("Constant", n)
		n.Constant.Walk(v)
		v.LeaveChildNode("Constant", n)
	}

	v.LeaveNode(n)
}
