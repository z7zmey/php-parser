package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Spaceship struct {
	BinaryOp
}

func NewSpaceship(variable node.Node, expression node.Node) node.Node {
	return Spaceship{
		BinaryOp{
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Spaceship) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Spaceship) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Spaceship) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Spaceship) Position() *node.Position {
	return n.position
}

func (n Spaceship) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Spaceship) Walk(v node.Visitor) {
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
