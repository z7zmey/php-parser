package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClassImplements node
type ClassImplements struct {
	InterfaceNames []node.Node
}

// NewClassImplements node constructor
func NewClassImplements(interfaceNames []node.Node) *ClassImplements {
	return &ClassImplements{
		interfaceNames,
	}
}

// Attributes returns node attributes as map
func (n *ClassImplements) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassImplements) Walk(v walker.Visitor) {
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
