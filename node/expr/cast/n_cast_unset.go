package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Unset node
type Unset struct {
	Expr node.Node
}

// NewUnset node constructor
func NewUnset(Expr node.Node) *Unset {
	return &Unset{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Unset) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Unset) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
