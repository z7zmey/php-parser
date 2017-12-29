package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Exit struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewExit(expr node.Node, isDie bool) node.Node {
	return Exit{
		"Exit",
		map[string]interface{}{
			"isDie": isDie,
		},
		expr,
	}
}

func (n Exit) Name() string {
	return "Exit"
}

func (n Exit) Attributes() map[string]interface{} {
	return n.attributes
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
