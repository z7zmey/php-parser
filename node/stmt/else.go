package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Else) Name() string {
	return "Else"
}

type Else struct {
	name  string
	token token.Token
	stmt  node.Node
}

func NewElse(token token.Token, stmt node.Node) node.Node {
	return Else{
		"Else",
		token,
		stmt,
	}
}

func (n Else) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.Children("stmt")
		n.stmt.Walk(vv)
	}
}
