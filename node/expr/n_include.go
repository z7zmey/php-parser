package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Include node
type Include struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewInclude node constructor
func NewInclude(Expression node.Node) *Include {
	return &Include{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Include) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Include) GetPosition() *position.Position {
	return n.Position
}

func (n *Include) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Include) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Include) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Include) Walk(v walker.Visitor) {
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
