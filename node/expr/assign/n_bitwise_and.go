package assign

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// BitwiseAnd node
type BitwiseAnd struct {
	Comments   []*comment.Comment
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewBitwiseAnd node constructor
func NewBitwiseAnd(Variable node.Node, Expression node.Node) *BitwiseAnd {
	return &BitwiseAnd{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *BitwiseAnd) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BitwiseAnd) GetPosition() *position.Position {
	return n.Position
}

func (n *BitwiseAnd) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *BitwiseAnd) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *BitwiseAnd) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseAnd) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expression != nil {
		v.EnterChildNode("Expression", n)
		n.Expression.Walk(v)
		v.LeaveChildNode("Expression", n)
	}

	v.LeaveNode(n)
}
