package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// List node
type List struct {
	Items []node.Node
}

// NewList node constructor
func NewList(Items []node.Node) *List {
	return &List{
		Items,
	}
}

// Attributes returns node attributes as map
func (n *List) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *List) Walk(v walker.Visitor) {
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
