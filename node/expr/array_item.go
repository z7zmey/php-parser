package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ArrayItem) Name() string {
	return "ArrayItem"
}

type ArrayItem struct {
	name  string
	key   node.Node
	val   node.Node
	byRef bool
}

func NewArrayItem(key node.Node, val node.Node, byRef bool) node.Node {
	return ArrayItem{
		"ArrayItem",
		key,
		val,
		byRef,
	}
}

func (n ArrayItem) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.key != nil {
		vv := v.Children("key")
		n.key.Walk(vv)
	}

	if n.val != nil {
		vv := v.Children("val")
		n.val.Walk(vv)
	}
}
