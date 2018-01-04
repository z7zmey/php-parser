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
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Mod) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Mod) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Mod) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Mod) Position() *node.Position {
	return n.position
}

func (n Mod) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
