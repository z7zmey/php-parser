package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// MethodCall node
type MethodCall struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variable     node.Node
	Method       node.Node
	ArgumentList *node.ArgumentList
}

// NewMethodCall node constructor
func NewMethodCall(Variable node.Node, Method node.Node, ArgumentList *node.ArgumentList) *MethodCall {
	return &MethodCall{
		FreeFloating: nil,
		Variable:     Variable,
		Method:       Method,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *MethodCall) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *MethodCall) GetPosition() *position.Position {
	return n.Position
}

func (n *MethodCall) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *MethodCall) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *MethodCall) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Method != nil {
		v.EnterChildNode("Method", n)
		n.Method.Walk(v)
		v.LeaveChildNode("Method", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	v.LeaveNode(n)
}
