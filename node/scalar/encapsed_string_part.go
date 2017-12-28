package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n EncapsedStringPart) Name() string {
	return "EncapsedStringPart"
}

type EncapsedStringPart struct {
	name  string
	token token.Token
}

func NewEncapsedStringPart(t token.Token) node.Node {
	return EncapsedStringPart{
		"EncapsedStringPart",
		t,
	}
}

func (n EncapsedStringPart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
