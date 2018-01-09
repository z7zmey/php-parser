package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type While struct {
	Token token.Token
	Cond  node.Node
	Stmt  node.Node
}

func NewWhile(Token token.Token, Cond node.Node, Stmt node.Node) *While {
	return &While{
		Token,
		Cond,
		Stmt,
	}
}

func (n *While) Attributes() map[string]interface{} {
	return nil
}

func (n *While) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
