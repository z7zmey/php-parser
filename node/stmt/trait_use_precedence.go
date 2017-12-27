package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

func (n TraitUsePrecedence) Name() string {
	return "TraitUsePrecedence"
}

type TraitUsePrecedence struct {
	name      string
	ref       node.Node
	insteadof node.Node
}

func NewTraitUsePrecedence(ref node.Node, insteadof node.Node) node.Node {
	return TraitUsePrecedence{
		"TraitUsePrecedence",
		ref,
		insteadof,
	}
}

func (n TraitUsePrecedence) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.ref != nil {
		vv := v.Children("ref")
		n.ref.Walk(vv)
	}

	if n.insteadof != nil {
		vv := v.Children("insteadof")
		n.insteadof.Walk(vv)
	}
}
