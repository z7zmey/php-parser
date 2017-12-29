package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Clone struct {
	name string
	expr node.Node
}

func NewClone(expression node.Node) node.Node {
	return Clone{
		"Clone",
		expression,
	}
}

func (n Clone) Name() string {
	return "Clone"
}

func (n Clone) Attributes() map[string]interface{} {
	return nil
}

func (n Clone) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
