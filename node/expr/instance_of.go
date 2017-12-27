package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n InstanceOf) Name() string {
	return "InstanceOf"
}

type InstanceOf struct {
	name  string
	expr  node.Node
	class node.Node
}

func NewInstanceOf(expr node.Node, class node.Node) node.Node {
	return InstanceOf{
		"InstanceOf",
		expr,
		class,
	}
}

func (n InstanceOf) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}

	if n.class != nil {
		vv := v.Children("class")
		n.class.Walk(vv)
	}
}
