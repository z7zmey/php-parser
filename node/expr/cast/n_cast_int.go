package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Int node
type Int struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewInt node constructor
func NewInt(Expr node.Node) *Int {
	return &Int{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Int) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Int) GetPosition() *position.Position {
	return n.Position
}

func (n *Int) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Int) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Int) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Int) Walk(v walker.Visitor) {
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
