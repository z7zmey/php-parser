package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Interface) Name() string {
	return "Interface"
}

type Interface struct {
	name          string
	token         token.Token
	interfaceName token.Token
	extends       []node.Node
	stmts         []node.Node
}

func NewInterface(token token.Token, name token.Token, extends []node.Node, stmts []node.Node) node.Node {
	return Interface{
		"Interface",
		token,
		name,
		extends,
		stmts,
	}
}

func (n Interface) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.interfaceName.Value)

	if n.extends != nil {
		vv := v.Children("extends")
		for _, nn := range n.extends {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
