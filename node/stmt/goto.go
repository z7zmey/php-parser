package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Goto) Name() string {
	return "Goto"
}

type Goto struct {
	name  string
	token token.Token
	label token.Token
}

// todl label must be identifier
func NewGoto(token token.Token, label token.Token) node.Node {
	return Goto{
		"Goto",
		token,
		label,
	}
}

func (n Goto) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("label", n.label.Value)
}
