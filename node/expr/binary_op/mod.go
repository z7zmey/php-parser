package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

func(n Mod) Name() string {
	return "Mod"
}

type Mod struct {
	BinaryOp
}

func NewMod(variable node.Node, expression node.Node) node.Node {
	return Mod{
		BinaryOp{
			"BinaryMod",
			variable,
			expression,
		},
	}
}
