package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Global) Name() string {
	return "Global"
}

type Global struct {
	name  string
	token token.Token
	vars  []node.Node
}

func NewGlobal(token token.Token, vars []node.Node) node.Node {
	return Global{
		"Global",
		token,
		vars,
	}
}

func (n Global) Walk(v node.Visitor) {
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
