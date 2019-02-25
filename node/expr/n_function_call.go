package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// FunctionCall node
type FunctionCall struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Function     node.Node
	ArgumentList *node.ArgumentList
}

// NewFunctionCall node constructor
func NewFunctionCall(Function node.Node, ArgumentList *node.ArgumentList) *FunctionCall {
	return &FunctionCall{
		FreeFloating: nil,
		Function:     Function,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *FunctionCall) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *FunctionCall) GetPosition() *position.Position {
	return n.Position
}

func (n *FunctionCall) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *FunctionCall) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *FunctionCall) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Function != nil {
		v.EnterChildNode("Function", n)
		n.Function.Walk(v)
		v.LeaveChildNode("Function", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	v.LeaveNode(n)
}
