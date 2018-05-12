package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// InterfaceExtends node
type InterfaceExtends struct {
	InterfaceNames []node.Node
}

// NewInterfaceExtends node constructor
func NewInterfaceExtends(InterfaceNames []node.Node) *InterfaceExtends {
	return &InterfaceExtends{
		InterfaceNames,
	}
}

// Attributes returns node attributes as map
func (n *InterfaceExtends) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InterfaceExtends) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InterfaceNames != nil {
		vv := v.GetChildrenVisitor("InterfaceNames")
		for _, nn := range n.InterfaceNames {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
