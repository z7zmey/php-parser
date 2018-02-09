package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// CastString node
type CastString struct {
	Expr node.Node
}

// NewCastString node constuctor
func NewCastString(Expr node.Node) *CastString {
	return &CastString{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *CastString) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastString) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
