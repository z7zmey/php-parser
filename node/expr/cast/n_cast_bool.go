package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Bool node
type Bool struct {
	Expr node.Node
}

// NewBool node constructor
func NewBool(Expr node.Node) *Bool {
	return &Bool{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Bool) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Bool) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
