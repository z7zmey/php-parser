package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PropertyFetch struct {
	Variable node.Node
	Property node.Node
}

func NewPropertyFetch(Variable node.Node, Property node.Node) *PropertyFetch {
	return &PropertyFetch{
		Variable,
		Property,
	}
}

func (n *PropertyFetch) Attributes() map[string]interface{} {
	return nil
}

func (n *PropertyFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Property != nil {
		vv := v.GetChildrenVisitor("Property")
		n.Property.Walk(vv)
	}

	v.LeaveNode(n)
}
