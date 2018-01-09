package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Continue struct {
	Expr node.Node
}

func NewContinue(Expr node.Node) *Continue {
	return &Continue{
		Expr,
	}
}

func (n *Continue) Attributes() map[string]interface{} {
	return nil
}

func (n *Continue) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
