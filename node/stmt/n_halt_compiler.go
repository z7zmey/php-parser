package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// HaltCompiler node
type HaltCompiler struct {
	Meta     []meta.Meta
	Position *position.Position
}

// NewHaltCompiler node constructor
func NewHaltCompiler() *HaltCompiler {
	return &HaltCompiler{}
}

// SetPosition sets node position
func (n *HaltCompiler) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *HaltCompiler) GetPosition() *position.Position {
	return n.Position
}

func (n *HaltCompiler) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *HaltCompiler) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *HaltCompiler) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
