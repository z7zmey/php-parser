package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Default) Name() string {
	return "Default"
}

type Default struct {
	name  string
	token token.Token
	stmts []node.Node
}

func NewDefault(token token.Token, stmts []node.Node) node.Node {
	return Default{
		"Default",
		token,
		stmts,
	}
}

func (n Default) Walk(v node.Visitor) {
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
