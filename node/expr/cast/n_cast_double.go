package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// CastDouble node
type CastDouble struct {
	Expr node.Node
}

// NewCastDouble node constuctor
func NewCastDouble(Expr node.Node) *CastDouble {
	return &CastDouble{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *CastDouble) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CastDouble) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
