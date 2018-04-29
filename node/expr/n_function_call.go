package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// FunctionCall node
type FunctionCall struct {
	Function     node.Node
	ArgumentList *node.ArgumentList
}

// NewFunctionCall node constructor
func NewFunctionCall(Function node.Node, ArgumentList *node.ArgumentList) *FunctionCall {
	return &FunctionCall{
		Function,
		ArgumentList,
	}
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
		vv := v.GetChildrenVisitor("Function")
		n.Function.Walk(vv)
	}

	if n.ArgumentList != nil {
		vv := v.GetChildrenVisitor("ArgumentList")
		n.ArgumentList.Walk(vv)
	}

	v.LeaveNode(n)
}
