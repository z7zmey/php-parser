package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Exit) Name() string {
	return "Exit"
}

type Exit struct {
	name  string
	expr  node.Node
	isDie bool
}

func NewExit(expr node.Node, isDie bool) node.Node {
	return Exit{
		"Exit",
		expr,
		isDie,
	}
}

func (n Exit) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
