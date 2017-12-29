package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Empty struct {
	name       string
	attributes map[string]interface{}
	expr       node.Node
}

func NewEmpty(expression node.Node) node.Node {
	return Empty{
		"Empty",
		map[string]interface{}{},
		expression,
	}
}

func (n Empty) Name() string {
	return "Empty"
}

func (n Empty) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Empty) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
