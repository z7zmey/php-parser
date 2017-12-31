package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Coalesce struct {
	BinaryOp
}

func NewCoalesce(variable node.Node, expression node.Node) node.Node {
	return Coalesce{
		BinaryOp{
			"BinaryCoalesce",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Coalesce) Name() string {
	return "Coalesce"
}

func (n Coalesce) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Coalesce) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Coalesce) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Coalesce) Position() *node.Position {
	return n.position
}

func (n Coalesce) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Coalesce) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
