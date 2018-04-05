package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Encapsed node
type Encapsed struct {
	Parts []node.Node
}

// NewEncapsed node constructor
func NewEncapsed(Parts []node.Node) *Encapsed {
	return &Encapsed{
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *Encapsed) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Encapsed) Walk(v walker.Visitor) {
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
}
