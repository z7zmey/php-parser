package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Array node
type Array struct {
	Items []node.Node
}

// NewArray node constructor
func NewArray(Items []node.Node) *Array {
	return &Array{
		Items,
	}
}

// Attributes returns node attributes as map
func (n *Array) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Array) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		vv := v.GetChildrenVisitor("Items")
		for _, nn := range n.Items {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
