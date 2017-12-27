package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Yield) Name() string {
	return "Yield"
}

type Yield struct {
	name  string
	key   node.Node
	value node.Node
}

func NewYield(key node.Node, value node.Node) node.Node {
	return Yield{
		"Yield",
		key,
		value,
	}
}

func (n Yield) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.key != nil {
		vv := v.Children("key")
		n.key.Walk(vv)
	}

	if n.value != nil {
		vv := v.Children("value")
		n.value.Walk(vv)
	}
}
