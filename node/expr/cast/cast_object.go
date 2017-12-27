package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastObject) Name() string {
	return "CastObject"
}

type CastObject struct {
	Cast
}

func NewCastObject(expr node.Node) node.Node {
	return CastObject{
		Cast{
			"CastObject",
			expr,
		},
	}
}

func (n CastObject) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
