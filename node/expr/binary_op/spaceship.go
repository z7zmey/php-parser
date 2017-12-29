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
			"BinarySpaceship",
			map[string]interface{}{},
			variable,
			expression,
		},
	}
}

func (n Spaceship) Name() string {
	return "Spaceship"
}

func (n Spaceship) Attributes() map[string]interface{} {
	return n.attributes
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
