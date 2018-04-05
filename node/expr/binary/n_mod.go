package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Mod node
type Mod struct {
	Left  node.Node
	Right node.Node
}

// NewMod node constructor
func NewMod(Variable node.Node, Expression node.Node) *Mod {
	return &Mod{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Mod) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Mod) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}
