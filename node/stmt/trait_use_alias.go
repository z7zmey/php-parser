package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n TraitUseAlias) Name() string {
	return "TraitUseAlias"
}

type TraitUseAlias struct {
	name     string
	ref      node.Node
	modifier node.Node
	alias    token.TokenInterface
}

func NewTraitUseAlias(ref node.Node, modifier node.Node, alias token.TokenInterface) node.Node {
	return TraitUseAlias{
		"TraitUseAlias",
		ref,
		modifier,
		alias,
	}
}

func (n TraitUseAlias) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.ref != nil {
		vv := v.Children("ref")
		n.ref.Walk(vv)
	}

	if n.modifier != nil {
		vv := v.Children("modifier")
		n.modifier.Walk(vv)
	}
}
