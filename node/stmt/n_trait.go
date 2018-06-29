package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Trait node
type Trait struct {
	Meta          []meta.Meta
	Position      *position.Position
	PhpDocComment string
	TraitName     node.Node
	Stmts         []node.Node
}

// NewTrait node constructor
func NewTrait(TraitName node.Node, Stmts []node.Node, PhpDocComment string) *Trait {
	return &Trait{
		PhpDocComment: PhpDocComment,
		TraitName:     TraitName,
		Stmts:         Stmts,
	}
}

// SetPosition sets node position
func (n *Trait) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Trait) GetPosition() *position.Position {
	return n.Position
}

func (n *Trait) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Trait) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Trait) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Trait) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.TraitName != nil {
		v.EnterChildNode("TraitName", n)
		n.TraitName.Walk(v)
		v.LeaveChildNode("TraitName", n)
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
