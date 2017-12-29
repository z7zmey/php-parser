package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ClusureUse struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
}

func NewClusureUse(variable node.Node, byRef bool) node.Node {
	return ClusureUse{
		"ClusureUse",
		map[string]interface{}{
			"byRef": byRef,
		},
		variable,
	}
}

func (n ClusureUse) Name() string {
	return "ClusureUse"
}

func (n ClusureUse) Attributes() map[string]interface{} {
	return nil
}

func (n ClusureUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
