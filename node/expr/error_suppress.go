package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ErrorSuppress struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return ErrorSuppress{
		"ErrorSuppress",
		map[string]interface{}{},
		expression,
	}
}

func (n ErrorSuppress) Name() string {
	return "ErrorSuppress"
}

func (n ErrorSuppress) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ErrorSuppress) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
