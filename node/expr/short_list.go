package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShortList) Name() string {
	return "ShortList"
}

type ShortList struct {
	name  string
	items []node.Node
}

func NewShortList(items []node.Node) node.Node {
	return ShortList{
		"ShortList",
		items,
	}
}

func (n ShortList) Walk(v node.Visitor) {
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
