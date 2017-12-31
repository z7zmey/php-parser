package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Assign struct {
	AssignOp
}

func NewAssign(variable node.Node, expression node.Node) node.Node {
	return Assign{
		AssignOp{
			"Assign",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Assign) Name() string {
	return "Assign"
}

func (n Assign) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Assign) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Assign) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Assign) Position() *node.Position {
	return n.position
}

func (n Assign) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Assign) Walk(v node.Visitor) {
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
