package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// LogicalXor node
type LogicalXor struct {
	Left  node.Node
	Right node.Node
}

// NewLogicalXor node constructor
func NewLogicalXor(Variable node.Node, Expression node.Node) *LogicalXor {
	return &LogicalXor{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *LogicalXor) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *LogicalXor) Walk(v walker.Visitor) {
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
