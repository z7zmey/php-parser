package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n String) Name() string {
	return "String"
}

type String struct {
	name  string
	token token.Token
}

func NewString(token token.Token) node.Node {
	return String{
		"String",
		token,
	}
}

func (n String) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
