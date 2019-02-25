package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// AltIf node
type AltIf struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Cond         node.Node
	Stmt         node.Node
	ElseIf       []node.Node
	Else         node.Node
}

// NewAltIf node constructor
func NewAltIf(Cond node.Node, Stmt node.Node, ElseIf []node.Node, Else node.Node) *AltIf {
	return &AltIf{
		FreeFloating: nil,
		Cond:         Cond,
		Stmt:         Stmt,
		ElseIf:       ElseIf,
		Else:         Else,
	}
}

// SetPosition sets node position
func (n *AltIf) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *AltIf) GetPosition() *position.Position {
	return n.Position
}

func (n *AltIf) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *AltIf) Attributes() map[string]interface{} {
	return nil
}

// AddElseIf add AltElseIf node and returns AltIf node
func (n *AltIf) AddElseIf(ElseIf node.Node) node.Node {
	if n.ElseIf == nil {
		n.ElseIf = make([]node.Node, 0)
	}

	n.ElseIf = append(n.ElseIf, ElseIf)

	return n
}

// SetElse set AltElse node and returns AltIf node
func (n *AltIf) SetElse(Else node.Node) node.Node {
	n.Else = Else

	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltIf) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	if n.ElseIf != nil {
		v.EnterChildList("ElseIf", n)
		for _, nn := range n.ElseIf {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("ElseIf", n)
	}

	if n.Else != nil {
		v.EnterChildNode("Else", n)
		n.Else.Walk(v)
		v.LeaveChildNode("Else", n)
	}

	v.LeaveNode(n)
}
