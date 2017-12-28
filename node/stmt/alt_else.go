package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n AltElse) Name() string {
	return "AltElse"
}

type AltElse struct {
	name  string
	token token.Token
	stmt  node.Node
}

func NewAltElse(token token.Token, stmt node.Node) node.Node {
	return AltElse{
		"AltElse",
		token,
		stmt,
	}
}

func (n AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
