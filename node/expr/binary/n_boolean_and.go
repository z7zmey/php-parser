package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// BooleanAnd node
type BooleanAnd struct {
	Left  node.Node
	Right node.Node
}

// NewBooleanAnd node constructor
func NewBooleanAnd(Variable node.Node, Expression node.Node) *BooleanAnd {
	return &BooleanAnd{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *BooleanAnd) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BooleanAnd) Walk(v walker.Visitor) {
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
