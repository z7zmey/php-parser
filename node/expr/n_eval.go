package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Eval node
type Eval struct {
	Expr node.Node
}

// NewEval node constructor
func NewEval(Expression node.Node) *Eval {
	return &Eval{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Eval) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Eval) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
