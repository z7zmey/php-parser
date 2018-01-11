package expr

import (
	"github.com/z7zmey/php-parser/node"
)

// ArrayDimFetch node
type ArrayDimFetch struct {
	Variable node.Node
	Dim      node.Node
}

// NewArrayDimFetch node constuctor
func NewArrayDimFetch(Variable node.Node, Dim node.Node) *ArrayDimFetch {
	return &ArrayDimFetch{
		Variable,
		Dim,
	}
}

// Attributes returns node attributes as map
func (n *ArrayDimFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
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
