package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n MagicConstant) Name() string {
	return "MagicConstant"
}

type MagicConstant struct {
	name  string
	token token.Token
}

func NewMagicConstant(token token.Token) node.Node {
	return MagicConstant{
		"MagicConstant",
		token,
	}
}

func (n MagicConstant) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
