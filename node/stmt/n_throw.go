package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Throw node
type Throw struct {
	Expr node.Node
}

// NewThrow node constructor
func NewThrow(Expr node.Node) *Throw {
	return &Throw{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Throw) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Throw) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
