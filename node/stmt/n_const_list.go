package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ConstList node
type ConstList struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Consts       []node.Node
}

// NewConstList node constructor
func NewConstList(Consts []node.Node) *ConstList {
	return &ConstList{
		FreeFloating: nil,
		Consts:       Consts,
	}
}

// SetPosition sets node position
func (n *ConstList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ConstList) GetPosition() *position.Position {
	return n.Position
}

func (n *ConstList) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ConstList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ConstList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		v.EnterChildList("Consts", n)
		for _, nn := range n.Consts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Consts", n)
	}

	v.LeaveNode(n)
}
