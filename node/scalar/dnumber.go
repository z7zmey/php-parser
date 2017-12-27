package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Dnumber) Name() string {
	return "Dnumber"
}

type Dnumber struct {
	name  string
	token token.Token
}

func NewDnumber(token token.Token) node.Node {
	return Dnumber{
		"Dnumber",
		token,
	}
}

func (n Dnumber) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
