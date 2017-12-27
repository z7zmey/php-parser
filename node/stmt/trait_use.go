package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n TraitUse) Name() string {
	return "TraitUse"
}

type TraitUse struct {
	name        string
	token       token.Token
	traits      []node.Node
	adaptations []node.Node
}

//TODO: traits myst be []node.Node
func NewTraitUse(token token.Token, traits []node.Node, adaptations []node.Node) node.Node {
	return TraitUse{
		"TraitUse",
		token,
		traits,
		adaptations,
	}
}

func (n TraitUse) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.traits != nil {
		vv := v.Children("traits")
		for _, nn := range n.traits {
			nn.Walk(vv)
		}
	}

	if n.adaptations != nil {
		vv := v.Children("adaptations")
		for _, nn := range n.adaptations {
			nn.Walk(vv)
		}
	}
}
