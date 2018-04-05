package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// YieldFrom node
type YieldFrom struct {
	Expr node.Node
}

// NewYieldFrom node constructor
func NewYieldFrom(Expression node.Node) *YieldFrom {
	return &YieldFrom{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *YieldFrom) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *YieldFrom) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
