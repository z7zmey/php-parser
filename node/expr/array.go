package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Array) Name() string {
	return "Array"
}

type Array struct {
	name       string
	opentToken token.Token
	closeToken token.Token
	items      []node.Node
}

func NewArray(opentToken token.Token, closeToken token.Token, items []node.Node) node.Node {
	return Array{
		"Array",
		opentToken,
		closeToken,
		items,
	}
}

func (n Array) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.items != nil {
		vv := v.GetChildrenVisitor("items")
		for _, nn := range n.items {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
