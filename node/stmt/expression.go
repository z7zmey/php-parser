package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	Expr node.Node
}

func NewExpression(Expr node.Node) *Expression {
	return &Expression{
		Expr,
	}
}

func (n *Expression) Attributes() map[string]interface{} {
	return nil
}

func (n *Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
