package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// PropertyFetch node
type PropertyFetch struct {
	Comments []*comment.Comment
	Position *position.Position
	Variable node.Node
	Property node.Node
}

// NewPropertyFetch node constructor
func NewPropertyFetch(Variable node.Node, Property node.Node) *PropertyFetch {
	return &PropertyFetch{
		Variable: Variable,
		Property: Property,
	}
}

// SetPosition sets node position
func (n *PropertyFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *PropertyFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *PropertyFetch) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *PropertyFetch) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *PropertyFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PropertyFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Property != nil {
		v.EnterChildNode("Property", n)
		n.Property.Walk(v)
		v.LeaveChildNode("Property", n)
	}

	v.LeaveNode(n)
}
