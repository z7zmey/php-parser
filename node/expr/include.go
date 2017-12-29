package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Include struct {
	name string
	expr node.Node
}

func NewInclude(expression node.Node) node.Node {
	return Include{
		"Include",
		expression,
	}
}

func (n Include) Name() string {
	return "Include"
}

func (n Include) Attributes() map[string]interface{} {
	return nil
}

func (n Include) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
