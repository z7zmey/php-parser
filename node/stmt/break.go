package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Break struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewBreak(expr node.Node) node.Node {
	return Break{
		"Break",
		map[string]interface{}{},
		expr,
	}
}

func (n Break) Name() string {
	return "Break"
}

func (n Break) Attributes() map[string]interface{} {
	return n.attributes
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
