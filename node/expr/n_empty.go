package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Empty node
type Empty struct {
	Position *position.Position
	Expr     node.Node
}

// NewEmpty node constructor
func NewEmpty(Expression node.Node) *Empty {
	return &Empty{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Empty) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Empty) GetPosition() *position.Position {
	return n.Position
}

// Attributes returns node attributes as map
func (n *Empty) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Empty) Walk(v walker.Visitor) {
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
