package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	AssignOp
}

func NewMod(variable  node.Node, expression node.Node) node.Node {
	return Mod{
		AssignOp{
			node.SimpleNode{Name: "AssignMod", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
