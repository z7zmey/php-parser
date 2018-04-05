package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Static node
type Static struct {
	Vars []node.Node
}

// NewStatic node constructor
func NewStatic(Vars []node.Node) *Static {
	return &Static{
		Vars,
	}
}

// Attributes returns node attributes as map
func (n *Static) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Static) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		vv := v.GetChildrenVisitor("Vars")
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
