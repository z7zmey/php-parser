package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// MethodCall node
type MethodCall struct {
	Variable  node.Node
	Method    node.Node
	Arguments []node.Node
}

// NewMethodCall node constructor
func NewMethodCall(Variable node.Node, Method node.Node, Arguments []node.Node) *MethodCall {
	return &MethodCall{
		Variable,
		Method,
		Arguments,
	}
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
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Method != nil {
		vv := v.GetChildrenVisitor("Method")
		n.Method.Walk(vv)
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
