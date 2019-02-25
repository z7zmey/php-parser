package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Foreach node
type Foreach struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Expr         node.Node
	Key          node.Node
	Variable     node.Node
	Stmt         node.Node
}

// NewForeach node constructor
func NewForeach(Expr node.Node, Key node.Node, Variable node.Node, Stmt node.Node) *Foreach {
	return &Foreach{
		FreeFloating: nil,
		Expr:         Expr,
		Key:          Key,
		Variable:     Variable,
		Stmt:         Stmt,
	}
}

// SetPosition sets node position
func (n *Foreach) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Foreach) GetPosition() *position.Position {
	return n.Position
}

func (n *Foreach) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Foreach) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Foreach) Walk(v walker.Visitor) {
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
