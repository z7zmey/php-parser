package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// AltForeach node
type AltForeach struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Expr         node.Node
	Key          node.Node
	Variable     node.Node
	Stmt         node.Node
}

// NewAltForeach node constructor
func NewAltForeach(Expr node.Node, Key node.Node, Variable node.Node, Stmt node.Node) *AltForeach {
	return &AltForeach{
		FreeFloating: nil,
		Expr:         Expr,
		Key:          Key,
		Variable:     Variable,
		Stmt:         Stmt,
	}
}

// SetPosition sets node position
func (n *AltForeach) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *AltForeach) GetPosition() *position.Position {
	return n.Position
}

func (n *AltForeach) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *AltForeach) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltForeach) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	if n.Key != nil {
		v.EnterChildNode("Key", n)
		n.Key.Walk(v)
		v.LeaveChildNode("Key", n)
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
