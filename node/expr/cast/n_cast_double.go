package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Double node
type Double struct {
	Expr node.Node
}

// NewDouble node constructor
func NewDouble(Expr node.Node) *Double {
	return &Double{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Double) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Double) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
