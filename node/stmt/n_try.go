package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Try node
type Try struct {
	Stmts   []node.Node
	Catches []node.Node
	Finally node.Node
}

// NewTry node constructor
func NewTry(Stmts []node.Node, Catches []node.Node, Finally node.Node) *Try {
	return &Try{
		Stmts,
		Catches,
		Finally,
	}
}

// Attributes returns node attributes as map
func (n *Try) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Try) Walk(v walker.Visitor) {
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

	if n.Catches != nil {
		vv := v.GetChildrenVisitor("Catches")
		for _, nn := range n.Catches {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Finally != nil {
		vv := v.GetChildrenVisitor("Finally")
		n.Finally.Walk(vv)
	}

	v.LeaveNode(n)
}
