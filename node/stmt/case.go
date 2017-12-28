package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Case) Name() string {
	return "Case"
}

type Case struct {
	name  string
	token token.Token
	cond  node.Node
	stmts []node.Node
}

func NewCase(token token.Token, cond node.Node, stmts []node.Node) node.Node {
	return Case{
		"Case",
		token,
		cond,
		stmts,
	}
}

func (n Case) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
