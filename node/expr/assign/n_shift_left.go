package assign

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ShiftLeft node
type ShiftLeft struct {
	Meta       []meta.Meta
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewShiftLeft node constructor
func NewShiftLeft(Variable node.Node, Expression node.Node) *ShiftLeft {
	return &ShiftLeft{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *ShiftLeft) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ShiftLeft) GetPosition() *position.Position {
	return n.Position
}

func (n *ShiftLeft) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ShiftLeft) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ShiftLeft) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShiftLeft) Walk(v walker.Visitor) {
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
