package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type StaticPropertyFetch struct {
	name     string
	class    node.Node
	property node.Node
}

func NewStaticPropertyFetch(class node.Node, property node.Node) node.Node {
	return StaticPropertyFetch{
		"StaticPropertyFetch",
		class,
		property,
	}
}

func (n StaticPropertyFetch) Name() string {
	return "StaticPropertyFetch"
}

func (n StaticPropertyFetch) Attributes() map[string]interface{} {
	return nil
}

func (n StaticPropertyFetch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.class != nil {
		vv := v.GetChildrenVisitor("class")
		n.class.Walk(vv)
	}

	if n.property != nil {
		vv := v.GetChildrenVisitor("property")
		n.property.Walk(vv)
	}

	v.LeaveNode(n)
}
