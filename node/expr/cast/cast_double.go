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

func (n CastDouble) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
