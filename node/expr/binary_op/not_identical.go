package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type NotIdentical struct {
	BinaryOp
}

func NewNotIdentical(variable node.Node, expression node.Node) node.Node {
	return NotIdentical{
		BinaryOp{
			"BinaryNotIdentical",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n NotIdentical) Name() string {
	return "NotIdentical"
}

func (n NotIdentical) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NotIdentical) Walk(v node.Visitor) {
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
