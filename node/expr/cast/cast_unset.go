package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastUnset struct {
	Cast
}

func NewCastUnset(expr node.Node) node.Node {
	return CastUnset{
		Cast{
			node.SimpleNode{Name: "CastUnset", Attributes: make(map[string]string)},
			expr,
		},
	}
}
