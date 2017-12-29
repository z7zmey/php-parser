package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUseAlias struct {
	name     string
	ref      node.Node
	modifier node.Node
	alias    node.Node
}

func NewTraitUseAlias(ref node.Node, modifier node.Node, alias node.Node) node.Node {
	return TraitUseAlias{
		"TraitUseAlias",
		ref,
		modifier,
		alias,
	}
}

func (n TraitUseAlias) Name() string {
	return "TraitUseAlias"
}

func (n TraitUseAlias) Attributes() map[string]interface{} {
	return nil
}

func (n TraitUseAlias) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ref != nil {
		vv := v.GetChildrenVisitor("ref")
		n.ref.Walk(vv)
	}

	if n.modifier != nil {
		vv := v.GetChildrenVisitor("modifier")
		n.modifier.Walk(vv)
	}

	if n.alias != nil {
		vv := v.GetChildrenVisitor("alias")
		n.alias.Walk(vv)
	}

	v.LeaveNode(n)
}
