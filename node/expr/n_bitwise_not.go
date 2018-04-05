package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// BitwiseNot node
type BitwiseNot struct {
	Expr node.Node
}

// NewBitwiseNot node constructor
func NewBitwiseNot(Expression node.Node) *BitwiseNot {
	return &BitwiseNot{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *BitwiseNot) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *BitwiseNot) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
