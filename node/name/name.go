package name

import (
	"github.com/z7zmey/php-parser/node"
)

// Name node
type Name struct {
	Parts []node.Node
}

// NewName node constuctor
func NewName(Parts []node.Node) *Name {
	return &Name{
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *Name) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Name) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
