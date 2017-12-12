package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastDouble struct {
	Cast
}

func NewCastDouble(expr node.Node) node.Node {
	return CastDouble{
		Cast{
			node.SimpleNode{Name: "CastDouble", Attributes: make(map[string]string)},
			expr,
		},
	}
}
