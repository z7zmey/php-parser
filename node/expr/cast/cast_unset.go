package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastUnset) Name() string {
	return "CastUnset"
}

type CastUnset struct {
	Cast
}

func NewCastUnset(expr node.Node) node.Node {
	return CastUnset{
		Cast{
			"CastUnset",
			expr,
		},
	}
}
