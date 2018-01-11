package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

// Div node
type Div struct {
	BinaryOp
}

// NewDiv node constuctor
func NewDiv(Variable node.Node, Expression node.Node) *Div {
	return &Div{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

// Attributes returns node attributes as map
func (n *Div) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Div) Walk(v node.Visitor) {
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
