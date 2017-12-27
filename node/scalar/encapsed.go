package scalar

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Encapsed) Name() string {
	return "Encapsed"
}

type Encapsed struct {
	name       string
	startToken token.Token
	endToken   token.Token
	parts      []node.Node
}

func NewEncapsed(startToken token.Token, parts []node.Node, endToken token.Token) node.Node {
	return Encapsed{
		"Encapsed",
		startToken,
		endToken,
		parts,
	}
}

func (n Encapsed) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.parts != nil {
		vv := v.Children("parts")
		for _, nn := range n.parts {
			nn.Walk(vv)
		}
	}
}
