package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Clone node
type Clone struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewClone node constructor
func NewClone(Expression node.Node) *Clone {
	return &Clone{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Clone) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Clone) GetPosition() *position.Position {
	return n.Position
}

func (n *Clone) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Clone) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Clone) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Clone) Walk(v walker.Visitor) {
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
