package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Throw) Name() string {
	return "Throw"
}

type Throw struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewThrow(token token.Token, expr node.Node) node.Node {
	return Throw{
		"Throw",
		token,
		expr,
	}
}

func (n Throw) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
