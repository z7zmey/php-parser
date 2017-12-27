package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n UnaryMinus) Name() string {
	return "UnaryMinus"
}

type UnaryMinus struct {
	name string
	expr node.Node
}

func NewUnaryMinus(expression node.Node) node.Node {
	return UnaryMinus{
		"UnaryMinus",
		expression,
	}
}

func (n UnaryMinus) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
