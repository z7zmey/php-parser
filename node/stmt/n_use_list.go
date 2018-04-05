package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// UseList node
type UseList struct {
	UseType node.Node
	Uses    []node.Node
}

// NewUseList node constructor
func NewUseList(UseType node.Node, Uses []node.Node) *UseList {
	return &UseList{
		UseType,
		Uses,
	}
}

// Attributes returns node attributes as map
func (n *UseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
