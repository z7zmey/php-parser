package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// UnaryPlus node
type UnaryPlus struct {
	Expr node.Node
}

// NewUnaryPlus node constructor
func NewUnaryPlus(Expression node.Node) *UnaryPlus {
	return &UnaryPlus{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *UnaryPlus) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UnaryPlus) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
