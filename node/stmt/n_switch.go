package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Switch node
type Switch struct {
	Cond  node.Node
	Cases []node.Node
}

// NewSwitch node constructor
func NewSwitch(Cond node.Node, Cases []node.Node) *Switch {
	return &Switch{
		Cond,
		Cases,
	}
}

// Attributes returns node attributes as map
func (n *Switch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Switch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Cases != nil {
		vv := v.GetChildrenVisitor("Cases")
		for _, nn := range n.Cases {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
