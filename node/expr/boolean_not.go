package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type BooleanNot struct {
	Expr node.Node
}

func NewBooleanNot(Expression node.Node) *BooleanNot {
	return &BooleanNot{
		Expression,
	}
}

func (n *BooleanNot) Attributes() map[string]interface{} {
	return nil
}

func (n *BooleanNot) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
