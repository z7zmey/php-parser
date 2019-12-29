package assign

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Coalesce node
type Coalesce struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variable     node.Node
	Expression   node.Node
}

// NewCoalesce node constructor
func NewCoalesce(Variable node.Node, Expression node.Node) *Coalesce {
	return &Coalesce{
		FreeFloating: nil,
		Variable:     Variable,
		Expression:   Expression,
	}
}

// SetPosition sets node position
func (n *Coalesce) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Coalesce) GetPosition() *position.Position {
	return n.Position
}

func (n *Coalesce) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Coalesce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Coalesce) Walk(v walker.Visitor) {
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
