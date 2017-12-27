package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Pow) Name() string {
	return "Pow"
}

type Pow struct {
	AssignOp
}

func NewPow(variable node.Node, expression node.Node) node.Node {
	return Pow{
		AssignOp{
			"AssignPow",
			variable,
			expression,
		},
	}
}
