package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Array node
type Array struct {
	Expr node.Node
}

// NewArray node constructor
func NewArray(Expr node.Node) *Array {
	return &Array{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Array) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Array) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
