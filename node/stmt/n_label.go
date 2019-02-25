package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Label node
type Label struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	LabelName    node.Node
}

// NewLabel node constructor
func NewLabel(LabelName node.Node) *Label {
	return &Label{
		FreeFloating: nil,
		LabelName:    LabelName,
	}
}

// SetPosition sets node position
func (n *Label) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Label) GetPosition() *position.Position {
	return n.Position
}

func (n *Label) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Label) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Label) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.LabelName != nil {
		v.EnterChildNode("LabelName", n)
		n.LabelName.Walk(v)
		v.LeaveChildNode("LabelName", n)
	}

	v.LeaveNode(n)
}
