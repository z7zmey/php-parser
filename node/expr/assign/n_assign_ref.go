package assign

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Reference node
type Reference struct {
	Meta       []meta.Meta
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewReference node constructor
func NewReference(Variable node.Node, Expression node.Node) *Reference {
	return &Reference{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *Reference) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Reference) GetPosition() *position.Position {
	return n.Position
}

func (n *Reference) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Reference) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Reference) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Reference) Walk(v walker.Visitor) {
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
