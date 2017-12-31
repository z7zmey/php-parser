package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	BinaryOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		BinaryOp{
			"BinaryMod",
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
