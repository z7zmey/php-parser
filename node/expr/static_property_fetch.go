package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n StaticPropertyFetch) Name() string {
	return "StaticPropertyFetch"
}

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

func (n StaticPropertyFetch) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.class != nil {
		vv := v.Children("class")
		n.class.Walk(vv)
	}

	if n.property != nil {
		vv := v.Children("property")
		n.property.Walk(vv)
	}
}
