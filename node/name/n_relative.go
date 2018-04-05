package name

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Relative node
type Relative struct {
	Parts []node.Node
}

// NewRelative node constructor
func NewRelative(Parts []node.Node) *Relative {
	return &Relative{
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *Relative) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Relative) Walk(v walker.Visitor) {
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

// GetParts returns the name parts
func (n *Relative) GetParts() []node.Node {
	return n.Parts
}
