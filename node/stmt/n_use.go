package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Use node
type Use struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	UseType      node.Node
	Use          node.Node
	Alias        node.Node
}

// NewUse node constructor
func NewUse(UseType node.Node, use node.Node, Alias node.Node) *Use {
	return &Use{
		FreeFloating: nil,
		UseType:      UseType,
		Use:          use,
		Alias:        Alias,
	}
}

// SetPosition sets node position
func (n *Use) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Use) GetPosition() *position.Position {
	return n.Position
}

func (n *Use) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Use) Attributes() map[string]interface{} {
	return nil
}

// SetUseType set use type and returns node
func (n *Use) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Use) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		v.EnterChildNode("UseType", n)
		n.UseType.Walk(v)
		v.LeaveChildNode("UseType", n)
	}

	if n.Use != nil {
		v.EnterChildNode("Use", n)
		n.Use.Walk(v)
		v.LeaveChildNode("Use", n)
	}

	if n.Alias != nil {
		v.EnterChildNode("Alias", n)
		n.Alias.Walk(v)
		v.LeaveChildNode("Alias", n)
	}

	v.LeaveNode(n)
}
