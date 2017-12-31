package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Pow struct {
	BinaryOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		BinaryOp{
			"BinaryPow",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n Pow) Name() string {
	return "Pow"
}

func (n Pow) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Pow) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Pow) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Pow) Position() *node.Position {
	return n.position
}

func (n Pow) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Pow) Walk(v node.Visitor) {
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
