package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Ternary) Name() string {
	return "Ternary"
}

type Ternary struct {
	name      string
	condition node.Node
	ifTrue    node.Node
	ifFalse   node.Node
}

func NewTernary(condition node.Node, ifTrue node.Node, ifFalse node.Node) node.Node {
	return Ternary{
		"Ternary",
		condition,
		ifTrue,
		ifFalse,
	}
}

func (n Ternary) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.condition != nil {
		vv := v.Children("condition")
		n.condition.Walk(vv)
	}

	if n.ifTrue != nil {
		vv := v.Children("ifTrue")
		n.ifTrue.Walk(vv)
	}

	if n.ifFalse != nil {
		vv := v.Children("ifFalse")
		n.ifFalse.Walk(vv)
	}
}
