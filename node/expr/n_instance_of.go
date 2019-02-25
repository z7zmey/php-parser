package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// InstanceOf node
type InstanceOf struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Expr         node.Node
	Class        node.Node
}

// NewInstanceOf node constructor
func NewInstanceOf(Expr node.Node, Class node.Node) *InstanceOf {
	return &InstanceOf{
		FreeFloating: nil,
		Expr:         Expr,
		Class:        Class,
	}
}

// SetPosition sets node position
func (n *InstanceOf) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *InstanceOf) GetPosition() *position.Position {
	return n.Position
}

func (n *InstanceOf) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *InstanceOf) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InstanceOf) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	if n.Class != nil {
		v.EnterChildNode("Class", n)
		n.Class.Walk(v)
		v.LeaveChildNode("Class", n)
	}

	v.LeaveNode(n)
}
