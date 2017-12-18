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
			node.SimpleNode{Name: "Assign", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
