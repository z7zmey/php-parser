package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Smaller node
type Smaller struct {
	Left  node.Node
	Right node.Node
}

// NewSmaller node constructor
func NewSmaller(Variable node.Node, Expression node.Node) *Smaller {
	return &Smaller{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Smaller) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Smaller) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
