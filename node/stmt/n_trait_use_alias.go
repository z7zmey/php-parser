package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// TraitUseAlias node
type TraitUseAlias struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Ref          node.Node
	Modifier     node.Node
	Alias        node.Node
}

// NewTraitUseAlias node constructor
func NewTraitUseAlias(Ref node.Node, Modifier node.Node, Alias node.Node) *TraitUseAlias {
	return &TraitUseAlias{
		FreeFloating: nil,
		Ref:          Ref,
		Modifier:     Modifier,
		Alias:        Alias,
	}
}

// SetPosition sets node position
func (n *TraitUseAlias) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *TraitUseAlias) GetPosition() *position.Position {
	return n.Position
}

func (n *TraitUseAlias) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *TraitUseAlias) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *TraitUseAlias) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Ref != nil {
		v.EnterChildNode("Ref", n)
		n.Ref.Walk(v)
		v.LeaveChildNode("Ref", n)
	}

	if n.Modifier != nil {
		v.EnterChildNode("Modifier", n)
		n.Modifier.Walk(v)
		v.LeaveChildNode("Modifier", n)
	}

	if n.Alias != nil {
		v.EnterChildNode("Alias", n)
		n.Alias.Walk(v)
		v.LeaveChildNode("Alias", n)
	}

	v.LeaveNode(n)
}
