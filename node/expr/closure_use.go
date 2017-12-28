package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ClusureUse) Name() string {
	return "ClusureUse"
}

type ClusureUse struct {
	name     string
	variable node.Node
	byRef    bool
}

func NewClusureUse(variable node.Node, byRef bool) node.Node {
	return ClusureUse{
		"ClusureUse",
		variable,
		byRef,
	}
}

func (n ClusureUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("byRef", n.byRef)

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
