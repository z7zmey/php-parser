package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Echo node
type Echo struct {
	Exprs []node.Node
}

// NewEcho node constructor
func NewEcho(Exprs []node.Node) *Echo {
	return &Echo{
		Exprs,
	}
}

// Attributes returns node attributes as map
func (n *Echo) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Echo) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		vv := v.GetChildrenVisitor("Exprs")
		for _, nn := range n.Exprs {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
