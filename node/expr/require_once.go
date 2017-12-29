package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type RequireOnce struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewRequireOnce(expression node.Node) node.Node {
	return RequireOnce{
		"RequireOnce",
		map[string]interface{}{},
		expression,
	}
}

func (n RequireOnce) Name() string {
	return "RequireOnce"
}

func (n RequireOnce) Attributes() map[string]interface{} {
	return n.attributes
}

func (n RequireOnce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
