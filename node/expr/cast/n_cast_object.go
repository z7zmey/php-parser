package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Object node
type Object struct {
	Comments []*comment.Comment
	Position *position.Position
	Expr     node.Node
}

// NewObject node constructor
func NewObject(Expr node.Node) *Object {
	return &Object{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Object) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Object) GetPosition() *position.Position {
	return n.Position
}

func (n *Object) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Object) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Object) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Object) Walk(v walker.Visitor) {
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
