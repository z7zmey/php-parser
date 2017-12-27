package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ShortArray) Name() string {
	return "ShortArray"
}

type ShortArray struct {
	name       string
	opentToken token.Token
	closeToken token.Token
	items      []node.Node
}

func NewShortArray(opentToken token.Token, closeToken token.Token, items []node.Node) node.Node {
	return ShortArray{
		"ShortArray",
		opentToken,
		closeToken,
		items,
	}
}

func (n ShortArray) Walk(v node.Visitor) {
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
