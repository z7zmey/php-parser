package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Property node
type Property struct {
	PhpDocComment string
	Variable      node.Node
	Expr          node.Node
}

// NewProperty node constructor
func NewProperty(Variable node.Node, Expr node.Node, PhpDocComment string) *Property {
	return &Property{
		PhpDocComment,
		Variable,
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Property) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Property) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
