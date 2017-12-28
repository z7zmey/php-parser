package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Do) Name() string {
	return "Do"
}

type Do struct {
	name  string
	token token.Token
	stmt  node.Node
	cond  node.Node
}

func NewDo(token token.Token, stmt node.Node, cond node.Node) node.Node {
	return Do{
		"Do",
		token,
		stmt,
		cond,
	}
}

func (n Do) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
