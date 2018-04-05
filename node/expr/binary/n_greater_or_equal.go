package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// GreaterOrEqual node
type GreaterOrEqual struct {
	Left  node.Node
	Right node.Node
}

// NewGreaterOrEqual node constructor
func NewGreaterOrEqual(Variable node.Node, Expression node.Node) *GreaterOrEqual {
	return &GreaterOrEqual{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *GreaterOrEqual) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *GreaterOrEqual) Walk(v walker.Visitor) {
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
