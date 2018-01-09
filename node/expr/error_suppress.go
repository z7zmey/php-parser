package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	Expr node.Node
}

func NewErrorSuppress(Expression node.Node) *ErrorSuppress {
	return &ErrorSuppress{
		Expression,
	}
}

func (n *ErrorSuppress) Attributes() map[string]interface{} {
	return nil
}

func (n *ErrorSuppress) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
