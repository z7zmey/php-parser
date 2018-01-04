package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Concat struct {
	AssignOp
}

func NewConcat(variable node.Node, expression node.Node) node.Node {
	return Concat{
		AssignOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Concat) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Concat) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Concat) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Concat) Position() *node.Position {
	return n.position
}

func (n Concat) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Concat) Walk(v node.Visitor) {
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
