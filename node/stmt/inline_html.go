package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n InlineHtml) Name() string {
	return "InlineHtml"
}

type InlineHtml struct {
	name  string
	token token.Token
}

func NewInlineHtml(token token.Token) node.Node {
	return InlineHtml{
		"InlineHtml",
		token,
	}
}

func (n InlineHtml) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	v.LeaveNode(n)
}
