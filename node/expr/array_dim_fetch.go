package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ArrayDimFetch) Name() string {
	return "ArrayDimFetch"
}

type ArrayDimFetch struct {
	name     string
	variable node.Node
	dim      node.Node
}

func NewArrayDimFetch(variable node.Node, dim node.Node) node.Node {
	return ArrayDimFetch{
		"ArrayDimFetch",
		variable,
		dim,
	}
}

func (n ArrayDimFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.dim != nil {
		vv := v.GetChildrenVisitor("dim")
		n.dim.Walk(vv)
	}

	v.LeaveNode(n)
}
