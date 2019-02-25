package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Else node
type Else struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Stmt         node.Node
}

// NewElse node constructor
func NewElse(Stmt node.Node) *Else {
	return &Else{
		FreeFloating: nil,
		Stmt:         Stmt,
	}
}

// SetPosition sets node position
func (n *Else) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Else) GetPosition() *position.Position {
	return n.Position
}

func (n *Else) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Else) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Else) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
