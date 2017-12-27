package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Trait) Name() string {
	return "Trait"
}

type Trait struct {
	name  string
	token token.Token
	stmts []node.Node
}

//TODO: stmts myst be []node.Node
func NewTrait(token token.Token, stmts []node.Node) node.Node {
	return Trait{
		"Trait",
		token,
		stmts,
	}
}

func (n Trait) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
