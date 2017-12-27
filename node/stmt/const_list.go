package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ConstList) Name() string {
	return "ConstList"
}

type ConstList struct {
	name   string
	token  token.Token
	consts []node.Node
}

func NewConstList(token token.Token, consts []node.Node) node.Node {
	return ConstList{
		"ConstList",
		token,
		consts,
	}
}

func (n ConstList) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.consts != nil {
		vv := v.Children("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}
}
