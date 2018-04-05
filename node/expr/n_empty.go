package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Empty node
type Empty struct {
	Expr node.Node
}

// NewEmpty node constructor
func NewEmpty(Expression node.Node) *Empty {
	return &Empty{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Empty) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Empty) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
