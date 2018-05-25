package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClosureUse node
type ClosureUse struct {
	Uses []node.Node
}

// NewClosureUse node constructor
func NewClosureUse(Uses []node.Node) *ClosureUse {
	return &ClosureUse{
		Uses,
	}
}

// Attributes returns node attributes as map
func (n *ClosureUse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClosureUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
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
