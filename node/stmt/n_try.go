package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Try node
type Try struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Stmts        []node.Node
	Catches      []node.Node
	Finally      node.Node
}

// NewTry node constructor
func NewTry(Stmts []node.Node, Catches []node.Node, Finally node.Node) *Try {
	return &Try{
		FreeFloating: nil,
		Stmts:        Stmts,
		Catches:      Catches,
		Finally:      Finally,
	}
}

// SetPosition sets node position
func (n *Try) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Try) GetPosition() *position.Position {
	return n.Position
}

func (n *Try) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Try) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Try) Walk(v walker.Visitor) {
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

	if n.Catches != nil {
		v.EnterChildList("Catches", n)
		for _, nn := range n.Catches {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Catches", n)
	}

	if n.Finally != nil {
		v.EnterChildNode("Finally", n)
		n.Finally.Walk(v)
		v.LeaveChildNode("Finally", n)
	}

	v.LeaveNode(n)
}
