package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Property node
type Property struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	PhpDocComment string
	Variable      node.Node
	Expr          node.Node
}

// NewProperty node constructor
func NewProperty(Variable node.Node, Expr node.Node, PhpDocComment string) *Property {
	return &Property{
		FreeFloating:  nil,
		PhpDocComment: PhpDocComment,
		Variable:      Variable,
		Expr:          Expr,
	}
}

// SetPosition sets node position
func (n *Property) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Property) GetPosition() *position.Position {
	return n.Position
}

func (n *Property) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Property) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Property) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
