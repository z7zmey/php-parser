package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Print) Name() string {
	return "Print"
}

type Print struct {
	name string
	expr node.Node
}

func NewPrint(expression node.Node) node.Node {
	return Print{
		"Print",
		expression,
	}
}

func (n Print) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
