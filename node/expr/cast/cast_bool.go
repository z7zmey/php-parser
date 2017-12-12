package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastBool struct {
	Cast
}

func NewCastBool(expr node.Node) node.Node {
	return CastBool{
		Cast{
			node.SimpleNode{Name: "CastBool", Attributes: make(map[string]string)},
			expr,
		},
	}
}
