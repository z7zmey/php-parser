package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PropertyFetch) Name() string {
	return "PropertyFetch"
}

type PropertyFetch struct {
	name     string
	variable node.Node
	property node.Node
}

func NewPropertyFetch(variable node.Node, property node.Node) node.Node {
	return PropertyFetch{
		"PropertyFetch",
		variable,
		property,
	}
}

func (n PropertyFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.property != nil {
		vv := v.GetChildrenVisitor("property")
		n.property.Walk(vv)
	}

	v.LeaveNode(n)
}
