package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// PreDec node
type PreDec struct {
	Meta     meta.Collection
	Position *position.Position
	Variable node.Node
}

// NewPreDec node constructor
func NewPreDec(Variable node.Node) *PreDec {
	return &PreDec{
		Variable: Variable,
	}
}

// SetPosition sets node position
func (n *PreDec) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *PreDec) GetPosition() *position.Position {
	return n.Position
}

func (n *PreDec) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *PreDec) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PreDec) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	v.LeaveNode(n)
}
