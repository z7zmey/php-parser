package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Mod) Name() string {
	return "Mod"
}

type Mod struct {
	AssignOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		AssignOp{
			"AssignMod",
			variable,
			expression,
		},
	}
}
