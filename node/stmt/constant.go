package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Constant) Name() string {
	return "Constant"
}

type Constant struct {
	name  string
	token token.Token
	expr  node.Node
}

func NewConstant(token token.Token, expr node.Node) node.Node {
	return Constant{
		"Constant",
		token,
		expr,
	}
}

func (n Constant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
