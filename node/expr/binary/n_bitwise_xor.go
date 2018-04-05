package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// BitwiseXor node
type BitwiseXor struct {
	Left  node.Node
	Right node.Node
}

// NewBitwiseXor node constructor
func NewBitwiseXor(Variable node.Node, Expression node.Node) *BitwiseXor {
	return &BitwiseXor{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *BitwiseXor) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseXor) Walk(v walker.Visitor) {
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
