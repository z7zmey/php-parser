package scalar

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Lnumber node
type Lnumber struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Value        string
}

// NewLnumber node constructor
func NewLnumber(Value string) *Lnumber {
	return &Lnumber{
		FreeFloating: nil,
		Value:        Value,
	}
}

// SetPosition sets node position
func (n *Lnumber) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Lnumber) GetPosition() *position.Position {
	return n.Position
}

func (n *Lnumber) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Lnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Lnumber) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
