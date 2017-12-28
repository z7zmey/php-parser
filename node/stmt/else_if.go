package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ElseIf) Name() string {
	return "ElseIf"
}

type ElseIf struct {
	name  string
	token token.Token
	cond  node.Node
	stmt  node.Node
}

func NewElseIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return ElseIf{
		"ElseIf",
		token,
		cond,
		stmt,
	}
}

func (n ElseIf) Walk(v node.Visitor) {
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
