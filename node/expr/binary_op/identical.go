package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

// Identical node
type Identical struct {
	BinaryOp
}

// NewIdentical node constuctor
func NewIdentical(Variable node.Node, Expression node.Node) *Identical {
	return &Identical{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

// Attributes returns node attributes as map
func (n *Identical) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Identical) Walk(v node.Visitor) {
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
