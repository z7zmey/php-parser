package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Else node
type Else struct {
	Comments []*comment.Comment
	Position *position.Position
	Stmt     node.Node
}

// NewElse node constructor
func NewElse(Stmt node.Node) *Else {
	return &Else{
		Stmt: Stmt,
	}
}

// SetPosition sets node position
func (n *Else) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Else) GetPosition() *position.Position {
	return n.Position
}

func (n *Else) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Else) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Else) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Else) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
