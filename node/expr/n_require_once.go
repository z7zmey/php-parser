package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// RequireOnce node
type RequireOnce struct {
	Expr node.Node
}

// NewRequireOnce node constructor
func NewRequireOnce(Expression node.Node) *RequireOnce {
	return &RequireOnce{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *RequireOnce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *RequireOnce) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
