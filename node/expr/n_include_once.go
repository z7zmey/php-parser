package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// IncludeOnce node
type IncludeOnce struct {
	Expr node.Node
}

// NewIncludeOnce node constructor
func NewIncludeOnce(Expression node.Node) *IncludeOnce {
	return &IncludeOnce{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *IncludeOnce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *IncludeOnce) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
