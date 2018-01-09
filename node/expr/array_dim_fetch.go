package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ArrayDimFetch struct {
	Variable node.Node
	Dim      node.Node
}

func NewArrayDimFetch(Variable node.Node, Dim node.Node) *ArrayDimFetch {
	return &ArrayDimFetch{
		Variable,
		Dim,
	}
}

func (n *ArrayDimFetch) Attributes() map[string]interface{} {
	return nil
}

func (n *ArrayDimFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Dim != nil {
		vv := v.GetChildrenVisitor("Dim")
		n.Dim.Walk(vv)
	}

	v.LeaveNode(n)
}
