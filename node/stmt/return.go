package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Return) Name() string {
	return "Return"
}

type Return struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewReturn(token token.Token, expr node.Node) node.Node {
	return Return{
		"Return",
		token,
		expr,
	}
}

func (n Return) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
