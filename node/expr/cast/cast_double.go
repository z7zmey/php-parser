package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastDouble) Name() string {
	return "CastDouble"
}

type CastDouble struct {
	Cast
}

func NewCastDouble(expr node.Node) node.Node {
	return CastDouble{
		Cast{
			"CastDouble",
			expr,
		},
	}
}
