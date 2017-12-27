package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Property) Name() string {
	return "Property"
}

type Property struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewProperty(token token.Token, expr node.Node) node.Node {
	return Property{
		"Property",
		token,
		expr,
	}
}

func (n Property) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
