package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Spaceship node
type Spaceship struct {
	Left  node.Node
	Right node.Node
}

// NewSpaceship node constructor
func NewSpaceship(Variable node.Node, Expression node.Node) *Spaceship {
	return &Spaceship{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Spaceship) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Spaceship) Walk(v walker.Visitor) {
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
