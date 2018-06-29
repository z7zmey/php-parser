package scalar

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Heredoc node
type Heredoc struct {
	Meta     []meta.Meta
	Position *position.Position
	Label    string
	Parts    []node.Node
}

// NewHeredoc node constructor
func NewHeredoc(Label string, Parts []node.Node) *Heredoc {
	return &Heredoc{
		Label: Label,
		Parts: Parts,
	}
}

// SetPosition sets node position
func (n *Heredoc) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Heredoc) GetPosition() *position.Position {
	return n.Position
}

func (n *Heredoc) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Heredoc) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Heredoc) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Label": n.Label,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Heredoc) Walk(v walker.Visitor) {
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
