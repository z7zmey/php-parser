package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// HaltCompiler node
type HaltCompiler struct {
	Comments []*comment.Comment
	Position *position.Position
}

// NewHaltCompiler node constructor
func NewHaltCompiler() *HaltCompiler {
	return &HaltCompiler{}
}

// SetPosition sets node position
func (n *HaltCompiler) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *HaltCompiler) GetPosition() *position.Position {
	return n.Position
}

func (n *HaltCompiler) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *HaltCompiler) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *HaltCompiler) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
