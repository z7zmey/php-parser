package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Label) Name() string {
	return "Label"
}

type Label struct {
	name  string
	token token.Token
}

func NewLabel(token token.Token) node.Node {
	return Label{
		"Label",
		token,
	}
}

func (n Label) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	v.LeaveNode(n)
}
