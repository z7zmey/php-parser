package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Finally) Name() string {
	return "Finally"
}

type Finally struct {
	name  string
	token token.Token
	stmts []node.Node
}

func NewFinally(token token.Token, stmts []node.Node) node.Node {
	return Finally{
		"Finally",
		token,
		stmts,
	}
}

func (n Finally) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
