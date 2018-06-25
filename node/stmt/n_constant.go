package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Constant node
type Constant struct {
	Comments      []*comment.Comment
	Position      *position.Position
	PhpDocComment string
	ConstantName  node.Node
	Expr          node.Node
}

// NewConstant node constructor
func NewConstant(ConstantName node.Node, Expr node.Node, PhpDocComment string) *Constant {
	return &Constant{
		PhpDocComment: PhpDocComment,
		ConstantName:  ConstantName,
		Expr:          Expr,
	}
}

// SetPosition sets node position
func (n *Constant) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Constant) GetPosition() *position.Position {
	return n.Position
}

func (n *Constant) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Constant) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Constant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Constant) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		v.EnterChildNode("ConstantName", n)
		n.ConstantName.Walk(v)
		v.LeaveChildNode("ConstantName", n)
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
