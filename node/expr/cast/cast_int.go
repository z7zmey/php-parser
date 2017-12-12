package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastInt struct {
	Cast
}

func NewCastInt(expr node.Node) node.Node {
	return CastInt{
		Cast{
			node.SimpleNode{Name: "CastInt", Attributes: make(map[string]string)},
			expr,
		},
	}
}
