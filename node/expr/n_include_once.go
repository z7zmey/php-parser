package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// IncludeOnce node
type IncludeOnce struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewIncludeOnce node constructor
func NewIncludeOnce(Expression node.Node) *IncludeOnce {
	return &IncludeOnce{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *IncludeOnce) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *IncludeOnce) GetPosition() *position.Position {
	return n.Position
}

func (n *IncludeOnce) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *IncludeOnce) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *IncludeOnce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *IncludeOnce) Walk(v walker.Visitor) {
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
