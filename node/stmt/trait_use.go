package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUse struct {
	name        string
	attributes  map[string]interface{}
	traits      []node.Node
	adaptations []node.Node
}

func NewTraitUse(traits []node.Node, adaptations []node.Node) node.Node {
	return TraitUse{
		"TraitUse",
		map[string]interface{}{},
		traits,
		adaptations,
	}
}

func (n TraitUse) Name() string {
	return "TraitUse"
}

func (n TraitUse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n TraitUse) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n TraitUse) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n TraitUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.traits != nil {
		vv := v.GetChildrenVisitor("traits")
		for _, nn := range n.traits {
			nn.Walk(vv)
		}
	}

	if n.adaptations != nil {
		vv := v.GetChildrenVisitor("adaptations")
		for _, nn := range n.adaptations {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
