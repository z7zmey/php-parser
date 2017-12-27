package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n TraitMethodRef) Name() string {
	return "TraitMethodRef"
}

type TraitMethodRef struct {
	name   string
	trait  node.Node
	method token.Token
}

// TODO: method must be identifier
func NewTraitMethodRef(trait node.Node, method token.Token) node.Node {
	return TraitMethodRef{
		"TraitMethodRef",
		trait,
		method,
	}
}

func (n TraitMethodRef) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("method", n.method.Value)

	if n.trait != nil {
		vv := v.Children("trait")
		n.trait.Walk(vv)
	}
}
