package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// StaticCall node
type StaticCall struct {
	Class     node.Node
	Call      node.Node
	Arguments []node.Node
}

// NewStaticCall node constructor
func NewStaticCall(Class node.Node, Call node.Node, Arguments []node.Node) *StaticCall {
	return &StaticCall{
		Class,
		Call,
		Arguments,
	}
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
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.Call != nil {
		vv := v.GetChildrenVisitor("Call")
		n.Call.Walk(vv)
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
