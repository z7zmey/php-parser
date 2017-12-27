package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Declare) Name() string {
	return "Declare"
}

type Declare struct {
	name   string
	token  token.Token
	consts []node.Node
	stmt   node.Node
}

func NewDeclare(token token.Token, consts []node.Node, stmt node.Node) node.Node {
	return Declare{
		"Declare",
		token,
		consts,
		stmt,
	}
}

func (n Declare) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.consts != nil {
		vv := v.Children("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	if n.stmt != nil {
		vv := v.Children("stmt")
		n.stmt.Walk(vv)
	}
}
