package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Include node
type Include struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewInclude node constructor
func NewInclude(Expression node.Node) *Include {
	return &Include{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Include) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Include) GetPosition() *position.Position {
	return n.Position
}

func (n *Include) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Include) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Include) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Include) Walk(v walker.Visitor) {
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
