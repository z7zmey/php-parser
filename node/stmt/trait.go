package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Trait struct {
	PhpDocComment string
	TraitName     node.Node
	Stmts         []node.Node
}

func NewTrait(TraitName node.Node, Stmts []node.Node, PhpDocComment string) *Trait {
	return &Trait{
		PhpDocComment,
		TraitName,
		Stmts,
	}
}

func (n *Trait) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n *Trait) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.TraitName != nil {
		vv := v.GetChildrenVisitor("TraitName")
		n.TraitName.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
