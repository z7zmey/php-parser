package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Catch) Name() string {
	return "Catch"
}

type Catch struct {
	name     string
	token    token.Token
	types    []node.Node
	variable node.Node
	stmts    []node.Node
}

func NewCatch(token token.Token, types []node.Node, variable node.Node, stmts []node.Node) node.Node {
	return Catch{
		"Catch",
		token,
		types,
		variable,
		stmts,
	}
}

func (n Catch) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.types != nil {
		vv := v.Children("types")
		for _, nn := range n.types {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
