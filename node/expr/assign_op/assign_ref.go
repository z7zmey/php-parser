package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignRef struct {
	AssignOp
}

func NewAssignRef(variable node.Node, expression node.Node) node.Node {
	return AssignRef{
		AssignOp{
			"AssignRef",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n AssignRef) Name() string {
	return "AssignRef"
}

func (n AssignRef) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AssignRef) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n AssignRef) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n AssignRef) Position() *node.Position {
	return n.position
}

func (n AssignRef) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n AssignRef) Walk(v node.Visitor) {
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
