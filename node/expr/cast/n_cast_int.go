package cast

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Int node
type Int struct {
	Expr node.Node
}

// NewInt node constructor
func NewInt(Expr node.Node) *Int {
	return &Int{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Int) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Int) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
