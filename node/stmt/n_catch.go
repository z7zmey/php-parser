package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Catch node
type Catch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Types        []node.Node
	Variable     node.Node
	Stmts        []node.Node
}

// NewCatch node constructor
func NewCatch(Types []node.Node, Variable node.Node, Stmts []node.Node) *Catch {
	return &Catch{
		FreeFloating: nil,
		Types:        Types,
		Variable:     Variable,
		Stmts:        Stmts,
	}
}

// SetPosition sets node position
func (n *Catch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Catch) GetPosition() *position.Position {
	return n.Position
}

func (n *Catch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Catch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Catch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Types != nil {
		v.EnterChildList("Types", n)
		for _, nn := range n.Types {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Types", n)
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
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
