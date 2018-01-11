package cast

import (
	"github.com/z7zmey/php-parser/node"
)

// CastObject node
type CastObject struct {
	Cast
}

// NewCastObject node constuctor
func NewCastObject(Expr node.Node) *CastObject {
	return &CastObject{
		Cast{
			Expr,
		},
	}
}

// Attributes returns node attributes as map
func (n *CastObject) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastObject) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
