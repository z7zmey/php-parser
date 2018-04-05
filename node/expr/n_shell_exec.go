package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ShellExec node
type ShellExec struct {
	Parts []node.Node
}

// NewShellExec node constructor
func NewShellExec(Parts []node.Node) *ShellExec {
	return &ShellExec{
		Parts,
	}
}

// Attributes returns node attributes as map
func (n *ShellExec) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShellExec) Walk(v walker.Visitor) {
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
