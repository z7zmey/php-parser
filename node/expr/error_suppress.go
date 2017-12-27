package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ErrorSuppress) Name() string {
	return "ErrorSuppress"
}

type ErrorSuppress struct {
	name string
	expr node.Node
}

func NewErrorSuppress(expression node.Node) node.Node {
	return ErrorSuppress{
		"ErrorSuppress",
		expression,
	}
}

func (n ErrorSuppress) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
