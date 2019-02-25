package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// AltSwitch node
type AltSwitch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Cond         node.Node
	CaseList     *CaseList
}

// NewAltSwitch node constructor
func NewAltSwitch(Cond node.Node, CaseList *CaseList) *AltSwitch {
	return &AltSwitch{
		FreeFloating: nil,
		Cond:         Cond,
		CaseList:     CaseList,
	}
}

// SetPosition sets node position
func (n *AltSwitch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *AltSwitch) GetPosition() *position.Position {
	return n.Position
}

func (n *AltSwitch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *AltSwitch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltSwitch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.CaseList != nil {
		v.EnterChildNode("CaseList", n)
		n.CaseList.Walk(v)
		v.LeaveChildNode("CaseList", n)
	}

	v.LeaveNode(n)
}
