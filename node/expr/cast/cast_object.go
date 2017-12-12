package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type CastObject struct {
	Cast
}

func NewCastObject(expr node.Node) node.Node {
	return CastObject{
		Cast{
			node.SimpleNode{Name: "CastObject", Attributes: make(map[string]string)},
			expr,
		},
	}
}
