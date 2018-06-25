package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Goto node
type Goto struct {
	Comments []*comment.Comment
	Position *position.Position
	Label    node.Node
}

// NewGoto node constructor
func NewGoto(Label node.Node) *Goto {
	return &Goto{
		Label: Label,
	}
}

// SetPosition sets node position
func (n *Goto) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Goto) GetPosition() *position.Position {
	return n.Position
}

func (n *Goto) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Goto) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Goto) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Goto) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Label != nil {
		v.EnterChildNode("Label", n)
		n.Label.Walk(v)
		v.LeaveChildNode("Label", n)
	}

	v.LeaveNode(n)
}
