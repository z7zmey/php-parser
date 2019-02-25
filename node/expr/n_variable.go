package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Variable node
type Variable struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	VarName      node.Node
}

// NewVariable node constructor
func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		FreeFloating: nil,
		VarName:      VarName,
	}
}

// SetPosition sets node position
func (n *Variable) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Variable) GetPosition() *position.Position {
	return n.Position
}

func (n *Variable) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

// SetVarName reset var name
func (n *Variable) SetVarName(VarName node.Node) {
	n.VarName = VarName
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Variable) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		v.EnterChildNode("VarName", n)
		n.VarName.Walk(v)
		v.LeaveChildNode("VarName", n)
	}

	v.LeaveNode(n)
}
