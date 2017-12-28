package name

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n NamePart) Name() string {
	return "NamePart"
}

type NamePart struct {
	name  string
	token token.Token
}

func NewNamePart(token token.Token) node.Node {
	return NamePart{
		"NamePart",
		token,
	}
}

func (n NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
