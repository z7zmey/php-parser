package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Empty) Name() string {
	return "Empty"
}

type Empty struct {
	name string
	expr node.Node
}

func NewEmpty(expression node.Node) node.Node {
	return Empty{
		"Empty",
		expression,
	}
}

func (n Empty) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
