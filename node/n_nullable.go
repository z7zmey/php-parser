package node

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Nullable node
type Nullable struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     Node
}

// NewNullable node constructor
func NewNullable(Expression Node) *Nullable {
	return &Nullable{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Nullable) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Nullable) GetPosition() *position.Position {
	return n.Position
}

func (n *Nullable) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Nullable) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Nullable) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Nullable) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
