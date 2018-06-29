package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ShellExec node
type ShellExec struct {
	Meta     []meta.Meta
	Position *position.Position
	Parts    []node.Node
}

// NewShellExec node constructor
func NewShellExec(Parts []node.Node) *ShellExec {
	return &ShellExec{
		Parts: Parts,
	}
}

// SetPosition sets node position
func (n *ShellExec) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ShellExec) GetPosition() *position.Position {
	return n.Position
}

func (n *ShellExec) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ShellExec) GetMeta() []meta.Meta {
	return n.Meta
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
		v.EnterChildList("Parts", n)
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Parts", n)
	}

	v.LeaveNode(n)
}
