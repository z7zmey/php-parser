package assign

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// BitwiseOr node
type BitwiseOr struct {
	Meta       []meta.Meta
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewBitwiseOr node constructor
func NewBitwiseOr(Variable node.Node, Expression node.Node) *BitwiseOr {
	return &BitwiseOr{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *BitwiseOr) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *BitwiseOr) GetPosition() *position.Position {
	return n.Position
}

func (n *BitwiseOr) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *BitwiseOr) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *BitwiseOr) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseOr) Walk(v walker.Visitor) {
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
