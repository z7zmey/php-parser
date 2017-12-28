package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Echo) Name() string {
	return "Echo"
}

type Echo struct {
	name  string
	token token.Token
	exprs []node.Node
}

func NewEcho(token token.Token, exprs []node.Node) node.Node {
	return Echo{
		"Echo",
		token,
		exprs,
	}
}

func (n Echo) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.exprs != nil {
		vv := v.GetChildrenVisitor("exprs")
		for _, nn := range n.exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
