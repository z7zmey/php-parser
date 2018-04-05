package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// StaticPropertyFetch node
type StaticPropertyFetch struct {
	Class    node.Node
	Property node.Node
}

// NewStaticPropertyFetch node constructor
func NewStaticPropertyFetch(Class node.Node, Property node.Node) *StaticPropertyFetch {
	return &StaticPropertyFetch{
		Class,
		Property,
	}
}

// Attributes returns node attributes as map
func (n *StaticPropertyFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticPropertyFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	if n.Property != nil {
		vv := v.GetChildrenVisitor("Property")
		n.Property.Walk(vv)
	}

	v.LeaveNode(n)
}
