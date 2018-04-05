package binary

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// NotEqual node
type NotEqual struct {
	Left  node.Node
	Right node.Node
}

// NewNotEqual node constructor
func NewNotEqual(Variable node.Node, Expression node.Node) *NotEqual {
	return &NotEqual{
		Variable,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *NotEqual) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *NotEqual) Walk(v walker.Visitor) {
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
