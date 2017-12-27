package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Continue) Name() string {
	return "Continue"
}

type Continue struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewContinue(token token.Token, expr node.Node) node.Node {
	return Continue{
		"Continue",
		token,
		expr,
	}
}

func (n Continue) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
