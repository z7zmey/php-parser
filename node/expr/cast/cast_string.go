package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastString struct {
	Cast
}

func NewCastString(expr node.Node) node.Node {
	return CastString{
		Cast{
			node.SimpleNode{Name: "CastString", Attributes: make(map[string]string)},
			expr,
		},
	}
}
