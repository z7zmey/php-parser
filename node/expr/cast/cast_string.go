package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastString) Name() string {
	return "CastString"
}

type CastString struct {
	Cast
}

func NewCastString(expr node.Node) node.Node {
	return CastString{
		Cast{
			"CastString",
			expr,
		},
	}
}
