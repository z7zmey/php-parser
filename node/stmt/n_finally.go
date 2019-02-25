package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Finally node
type Finally struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Stmts        []node.Node
}

// NewFinally node constructor
func NewFinally(Stmts []node.Node) *Finally {
	return &Finally{
		FreeFloating: nil,
		Stmts:        Stmts,
	}
}

// SetPosition sets node position
func (n *Finally) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Finally) GetPosition() *position.Position {
	return n.Position
}

func (n *Finally) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Finally) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Finally) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
