package cast

import (
	"github.com/z7zmey/php-parser/node"
)

// CastString node
type CastString struct {
	Cast
}

// NewCastString node constuctor
func NewCastString(Expr node.Node) *CastString {
	return &CastString{
		Cast{
			Expr,
		},
	}
}

// Attributes returns node attributes as map
func (n *CastString) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastString) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
