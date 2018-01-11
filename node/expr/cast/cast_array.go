package cast

import (
	"github.com/z7zmey/php-parser/node"
)

// CastArray node
type CastArray struct {
	Cast
}

// NewCastArray node constuctor
func NewCastArray(Expr node.Node) *CastArray {
	return &CastArray{
		Cast{
			Expr,
		},
	}
}

// Attributes returns node attributes as map
func (n *CastArray) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastArray) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
