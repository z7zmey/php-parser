package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Break) Name() string {
	return "Break"
}

type Break struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewBreak(token token.Token, expr node.Node) node.Node {
	return Break{
		"Break",
		token,
		expr,
	}
}

func (n Break) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
