package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Clone node
type Clone struct {
	Expr node.Node
}

// NewClone node constructor
func NewClone(Expression node.Node) *Clone {
	return &Clone{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Clone) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Clone) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
