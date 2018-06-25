package binary

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// BooleanAnd node
type BooleanAnd struct {
	Comments []*comment.Comment
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewBooleanAnd node constructor
func NewBooleanAnd(Variable node.Node, Expression node.Node) *BooleanAnd {
	return &BooleanAnd{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *BooleanAnd) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BooleanAnd) GetPosition() *position.Position {
	return n.Position
}

func (n *BooleanAnd) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *BooleanAnd) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *BooleanAnd) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BooleanAnd) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
