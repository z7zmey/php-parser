package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type AltElse struct {
	Else
}

func NewAltElse(token token.Token, stmt node.Node) node.Node {
	return AltElse{
		Else{
			node.SimpleNode{Name: "AltElse", Attributes: make(map[string]string)},
			token,
			stmt,
		},
	}
}
