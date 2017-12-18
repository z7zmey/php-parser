package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type AltElseIf struct {
	ElseIf
}

func NewAltElseIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltElseIf{
		ElseIf{
			node.SimpleNode{Name: "AltElseIf", Attributes: make(map[string]string)},
			token,
			cond,
			stmt,
		},
	}
}
