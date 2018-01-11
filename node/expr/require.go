package expr

import (
	"github.com/z7zmey/php-parser/node"
)

// Require node
type Require struct {
	Expr node.Node
}

// NewRequire node constuctor
func NewRequire(Expression node.Node) *Require {
	return &Require{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Require) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Require) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
