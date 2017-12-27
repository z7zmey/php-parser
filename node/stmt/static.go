package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Static) Name() string {
	return "Static"
}

type Static struct {
	name  string
	token token.Token
	vars  []node.Node
}

func NewStatic(token token.Token, vars []node.Node) node.Node {
	return Static{
		"Static",
		token,
		vars,
	}
}

func (n Static) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.Children("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}
}
