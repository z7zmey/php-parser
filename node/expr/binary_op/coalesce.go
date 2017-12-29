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
