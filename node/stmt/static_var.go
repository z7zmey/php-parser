package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n StaticVar) Name() string {
	return "StaticVar"
}

type StaticVar struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewStaticVar(token token.Token, expr node.Node) node.Node {
	return StaticVar{
		"StaticVar",
		token,
		expr,
	}
}

func (n StaticVar) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
