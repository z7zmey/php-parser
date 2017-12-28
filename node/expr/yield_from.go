package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n YieldFrom) Name() string {
	return "YieldFrom"
}

type YieldFrom struct {
	name string
	expr node.Node
}

func NewYieldFrom(expression node.Node) node.Node {
	return YieldFrom{
		"YieldFrom",
		expression,
	}
}

func (n YieldFrom) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
