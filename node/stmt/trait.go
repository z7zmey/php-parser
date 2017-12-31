package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Trait struct {
	name       string
	attributes map[string]interface{}
	traitName  node.Node
	stmts      []node.Node
}

func NewTrait(traitName node.Node, stmts []node.Node) node.Node {
	return Trait{
		"Trait",
		map[string]interface{}{},
		traitName,
		stmts,
	}
}

func (n Trait) Name() string {
	return "Trait"
}

func (n Trait) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Trait) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Trait) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Trait) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.traitName != nil {
		vv := v.GetChildrenVisitor("traitName")
		n.traitName.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
