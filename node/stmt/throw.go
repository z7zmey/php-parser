package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Throw struct {
	Expr node.Node
}

func NewThrow(Expr node.Node) *Throw {
	return &Throw{
		Expr,
	}
}

func (n *Throw) Attributes() map[string]interface{} {
	return nil
}

func (n *Throw) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
