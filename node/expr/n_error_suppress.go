package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ErrorSuppress node
type ErrorSuppress struct {
	Expr node.Node
}

// NewErrorSuppress node constructor
func NewErrorSuppress(Expression node.Node) *ErrorSuppress {
	return &ErrorSuppress{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *ErrorSuppress) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ErrorSuppress) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
