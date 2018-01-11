package expr

import (
	"github.com/z7zmey/php-parser/node"
)

// List node
type List struct {
	Items []node.Node
}

// NewList node constuctor
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
func (n *List) Walk(v node.Visitor) {
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
