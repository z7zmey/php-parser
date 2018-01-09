package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Property struct {
	PhpDocComment string
	Variable      node.Node
	Expr          node.Node
}

func NewProperty(Variable node.Node, Expr node.Node, PhpDocComment string) *Property {
	return &Property{
		PhpDocComment,
		Variable,
		Expr,
	}
}
func (n *Property) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n *Property) Walk(v node.Visitor) {
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
