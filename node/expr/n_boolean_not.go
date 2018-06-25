package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// BooleanNot node
type BooleanNot struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewBooleanNot node constructor
func NewBooleanNot(Expression node.Node) *BooleanNot {
	return &BooleanNot{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *BooleanNot) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BooleanNot) GetPosition() *position.Position {
	return n.Position
}

func (n *BooleanNot) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *BooleanNot) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *BooleanNot) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BooleanNot) Walk(v walker.Visitor) {
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
