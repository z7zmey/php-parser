package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitMethodRef struct {
	name   string
	trait  node.Node
	method node.Node
}

func NewTraitMethodRef(trait node.Node, method node.Node) node.Node {
	return TraitMethodRef{
		"TraitMethodRef",
		trait,
		method,
	}
}

func (n TraitMethodRef) Name() string {
	return "TraitMethodRef"
}

func (n TraitMethodRef) Attributes() map[string]interface{} {
	return nil
}

func (n TraitMethodRef) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.trait != nil {
		vv := v.GetChildrenVisitor("trait")
		n.trait.Walk(vv)
	}

	if n.method != nil {
		vv := v.GetChildrenVisitor("method")
		n.method.Walk(vv)
	}

	v.LeaveNode(n)
}
