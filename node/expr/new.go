package expr

import (
	"github.com/z7zmey/php-parser/node"
)

// New node
type New struct {
	Class     node.Node
	Arguments []node.Node
}

// NewNew node constuctor
func NewNew(Class node.Node, Arguments []node.Node) *New {
	return &New{
		Class,
		Arguments,
	}
}

// Attributes returns node attributes as map
func (n *New) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *New) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
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
