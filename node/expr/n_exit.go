package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Exit node
type Exit struct {
	FreeFloating freefloating.Collection
	Die          bool
	Position     *position.Position
	Expr         node.Node
}

// NewExit node constructor
func NewExit(Expr node.Node) *Exit {
	return &Exit{
		FreeFloating: nil,
		Expr:         Expr,
	}
}

// SetPosition sets node position
func (n *Exit) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Exit) GetPosition() *position.Position {
	return n.Position
}

func (n *Exit) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Exit) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Die": n.Die,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Exit) Walk(v walker.Visitor) {
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
