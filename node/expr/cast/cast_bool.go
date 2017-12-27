package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastBool) Name() string {
	return "CastBool"
}

type CastBool struct {
	Cast
}

func NewCastBool(expr node.Node) node.Node {
	return CastBool{
		Cast{
			"CastBool",
			expr,
		},
	}
}
