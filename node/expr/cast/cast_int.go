package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastInt) Name() string {
	return "CastInt"
}

type CastInt struct {
	Cast
}

func NewCastInt(expr node.Node) node.Node {
	return CastInt{
		Cast{
			"CastInt",
			expr,
		},
	}
}
