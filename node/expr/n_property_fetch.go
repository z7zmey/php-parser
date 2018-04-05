package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PropertyFetch node
type PropertyFetch struct {
	Variable node.Node
	Property node.Node
}

// NewPropertyFetch node constructor
func NewPropertyFetch(Variable node.Node, Property node.Node) *PropertyFetch {
	return &PropertyFetch{
		Variable,
		Property,
	}
}

// Attributes returns node attributes as map
func (n *PropertyFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PropertyFetch) Walk(v walker.Visitor) {
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
