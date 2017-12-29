package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Eval struct {
	name string
	expr node.Node
}

func NewEval(expression node.Node) node.Node {
	return Eval{
		"Eval",
		expression,
	}
}

func (n Eval) Name() string {
	return "Eval"
}

func (n Eval) Attributes() map[string]interface{} {
	return nil
}

func (n Eval) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
