package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// StaticCall node
type StaticCall struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Class        node.Node
	Call         node.Node
	ArgumentList *node.ArgumentList
}

// NewStaticCall node constructor
func NewStaticCall(Class node.Node, Call node.Node, ArgumentList *node.ArgumentList) *StaticCall {
	return &StaticCall{
		FreeFloating: nil,
		Class:        Class,
		Call:         Call,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *StaticCall) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *StaticCall) GetPosition() *position.Position {
	return n.Position
}

func (n *StaticCall) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *StaticCall) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticCall) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		v.EnterChildNode("Class", n)
		n.Class.Walk(v)
		v.LeaveChildNode("Class", n)
	}

	if n.Call != nil {
		v.EnterChildNode("Call", n)
		n.Call.Walk(v)
		v.LeaveChildNode("Call", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	v.LeaveNode(n)
}
