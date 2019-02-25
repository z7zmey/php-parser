package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// GroupUse node
type GroupUse struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	UseType      node.Node
	Prefix       node.Node
	UseList      []node.Node
}

// NewGroupUse node constructor
func NewGroupUse(UseType node.Node, Prefix node.Node, UseList []node.Node) *GroupUse {
	return &GroupUse{
		FreeFloating: nil,
		UseType:      UseType,
		Prefix:       Prefix,
		UseList:      UseList,
	}
}

// SetPosition sets node position
func (n *GroupUse) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *GroupUse) GetPosition() *position.Position {
	return n.Position
}

func (n *GroupUse) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *GroupUse) Attributes() map[string]interface{} {
	return nil
}

// SetUseType set use type and returns node
func (n *GroupUse) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *GroupUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		v.EnterChildNode("UseType", n)
		n.UseType.Walk(v)
		v.LeaveChildNode("UseType", n)
	}

	if n.Prefix != nil {
		v.EnterChildNode("Prefix", n)
		n.Prefix.Walk(v)
		v.LeaveChildNode("Prefix", n)
	}

	if n.UseList != nil {
		v.EnterChildList("UseList", n)
		for _, nn := range n.UseList {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("UseList", n)
	}

	v.LeaveNode(n)
}
