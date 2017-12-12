package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mod struct {
	BinaryOp
}

func NewMod(variable  node.Node, expression node.Node) node.Node {
	return Mod{
		BinaryOp{
			node.SimpleNode{Name: "BinaryMod", Attributes: make(map[string]string)},
			variable,
			expression,
		},
	}
}
