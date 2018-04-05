package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Exit node
type Exit struct {
	Expr node.Node
}

// NewExit node constructor
func NewExit(Expr node.Node) *Exit {
	return &Exit{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Exit) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Exit) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
