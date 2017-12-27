package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Clone) Name() string {
	return "Clone"
}

type Clone struct {
	name string
	expr node.Node
}

func NewClone(expression node.Node) node.Node {
	return Clone{
		"Clone",
		expression,
	}
}

func (n Clone) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
