package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Expression node
type Expression struct {
	Expr node.Node
}

// NewExpression node constructor
func NewExpression(Expr node.Node) *Expression {
	return &Expression{
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Expression) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Expression) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
