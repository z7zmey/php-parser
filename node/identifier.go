package node

import (
	"github.com/z7zmey/php-parser/token"
)

type Identifier struct {
	name  string
	token token.Token
}

func (n Identifier) Name() string {
	return "Identifier"
}

func NewIdentifier(token token.Token) Node {
	return Identifier{
		"Identifier",
		token,
	}
}

func (n Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	v.LeaveNode(n)
}
