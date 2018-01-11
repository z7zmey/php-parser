package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

// Default node
type Default struct {
	Stmts []node.Node
}

// NewDefault node constuctor
func NewDefault(Stmts []node.Node) *Default {
	return &Default{
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *Default) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Default) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
