package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n IncludeOnce) Name() string {
	return "IncludeOnce"
}

type IncludeOnce struct {
	name string
	expr node.Node
}

func NewIncludeOnce(expression node.Node) node.Node {
	return IncludeOnce{
		"IncludeOnce",
		expression,
	}
}

func (n IncludeOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
