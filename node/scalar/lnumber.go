package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Lnumber) Name() string {
	return "Lnumber"
}

type Lnumber struct {
	name  string
	token token.Token
}

func NewLnumber(token token.Token) node.Node {
	return Lnumber{
		"Lnumber",
		token,
	}
}

func (n Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
}
