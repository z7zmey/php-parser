package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

// Break node
type Break struct {
	Expr node.Node
}

// NewBreak node constuctor
func NewBreak(Expr node.Node) *Break {
	return &Break{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Break) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Break) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
