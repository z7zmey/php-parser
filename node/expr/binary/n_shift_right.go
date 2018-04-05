package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ShiftRight node
type ShiftRight struct {
	Left  node.Node
	Right node.Node
}

// NewShiftRight node constructor
func NewShiftRight(Variable node.Node, Expression node.Node) *ShiftRight {
	return &ShiftRight{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *ShiftRight) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ShiftRight) Walk(v walker.Visitor) {
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
