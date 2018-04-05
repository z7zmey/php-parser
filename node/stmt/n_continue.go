package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Continue node
type Continue struct {
	Expr node.Node
}

// NewContinue node constructor
func NewContinue(Expr node.Node) *Continue {
	return &Continue{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Continue) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Continue) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
