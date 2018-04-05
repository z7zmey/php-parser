package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// FunctionCall node
type FunctionCall struct {
	Function  node.Node
	Arguments []node.Node
}

// NewFunctionCall node constructor
func NewFunctionCall(Function node.Node, Arguments []node.Node) *FunctionCall {
	return &FunctionCall{
		Function,
		Arguments,
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

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
