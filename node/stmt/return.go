package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Return struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewReturn(expr node.Node) node.Node {
	return Return{
		"Return",
		map[string]interface{}{},
		expr,
	}
}

func (n Return) Name() string {
	return "Return"
}

func (n Return) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
