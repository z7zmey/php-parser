package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n List) Name() string {
	return "List"
}

type List struct {
	name  string
	items []node.Node
}

func NewList(items []node.Node) node.Node {
	return List{
		"List",
		items,
	}
}

func (n List) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.items != nil {
		vv := v.Children("items")
		for _, nn := range n.items {
			nn.Walk(vv)
		}
	}
}
