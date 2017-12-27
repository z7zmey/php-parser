package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Assign) Name() string {
	return "Assign"
}

type Assign struct {
	AssignOp
}

func NewAssign(variable node.Node, expression node.Node) node.Node {
	return Assign{
		AssignOp{
			"Assign",
			variable,
			expression,
		},
	}
}
