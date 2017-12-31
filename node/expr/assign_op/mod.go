package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	AssignOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		AssignOp{
			"AssignMod",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Mod) Name() string {
	return "Mod"
}

func (n Mod) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Mod) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Mod) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Mod) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.expression != nil {
		vv := v.GetChildrenVisitor("expression")
		n.expression.Walk(vv)
	}

	v.LeaveNode(n)
}
