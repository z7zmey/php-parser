package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticVar struct {
	Variable node.Node
	Expr     node.Node
}

func NewStaticVar(Variable node.Node, Expr node.Node) *StaticVar {
	return &StaticVar{
		Variable,
		Expr,
	}
}

func (n *StaticVar) Attributes() map[string]interface{} {
	return nil
}

func (n *StaticVar) Walk(v node.Visitor) {
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
