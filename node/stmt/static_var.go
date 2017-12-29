package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticVar struct {
	name     string
	variable node.Node
	expr     node.Node
}

func NewStaticVar(variable node.Node, expr node.Node) node.Node {
	return StaticVar{
		"StaticVar",
		variable,
		expr,
	}
}

func (n StaticVar) Name() string {
	return "StaticVar"
}

func (n StaticVar) Attributes() map[string]interface{} {
	return nil
}

func (n StaticVar) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
