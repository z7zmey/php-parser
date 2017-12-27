package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Switch) Name() string {
	return "Switch"
}

type Switch struct {
	name  string
	token token.Token
	cond  node.Node
	cases []node.Node
}

func NewSwitch(token token.Token, cond node.Node, cases []node.Node) node.Node {
	return Switch{
		"Switch",
		token,
		cond,
		cases,
	}
}

func (n Switch) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.Children("cond")
		n.cond.Walk(vv)
	}

	if n.cases != nil {
		vv := v.Children("cases")
		for _, nn := range n.cases {
			nn.Walk(vv)
		}
	}
}
