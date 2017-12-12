package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastArray struct {
	Cast
}

func NewCastArray(expr node.Node) node.Node {
	return CastArray{
		Cast{
			node.SimpleNode{Name: "CastArray", Attributes: make(map[string]string)},
			expr,
		},
	}
}
