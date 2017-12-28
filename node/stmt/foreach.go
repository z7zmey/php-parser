package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Foreach) Name() string {
	return "Foreach"
}

type Foreach struct {
	name     string
	token    token.Token
	expr     node.Node
	key      node.Node
	variable node.Node
	stmt     node.Node
	byRef    bool
}

func NewForeach(token token.Token, expr node.Node, key node.Node, variable node.Node, stmt node.Node, byRef bool) node.Node {
	return Foreach{
		"Foreach",
		token,
		expr,
		key,
		variable,
		stmt,
		byRef,
	}
}

func (n Foreach) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	if n.key != nil {
		vv := v.GetChildrenVisitor("key")
		n.key.Walk(vv)
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
