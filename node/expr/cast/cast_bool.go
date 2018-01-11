package cast

import (
	"github.com/z7zmey/php-parser/node"
)

// CastBool node
type CastBool struct {
	Cast
}

// NewCastBool node constuctor
func NewCastBool(Expr node.Node) *CastBool {
	return &CastBool{
		Cast{
			Expr,
		},
	}
}

// Attributes returns node attributes as map
func (n *CastBool) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastBool) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
