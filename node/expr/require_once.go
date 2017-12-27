package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n RequireOnce) Name() string {
	return "RequireOnce"
}

type RequireOnce struct {
	name string
	expr node.Node
}

func NewRequireOnce(expression node.Node) node.Node {
	return RequireOnce{
		"RequireOnce",
		expression,
	}
}

func (n RequireOnce) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
