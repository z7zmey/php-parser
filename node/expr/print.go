package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Print struct {
	Expr node.Node
}

func NewPrint(Expression node.Node) *Print {
	return &Print{
		Expression,
	}
}

func (n *Print) Attributes() map[string]interface{} {
	return nil
}

func (n *Print) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
