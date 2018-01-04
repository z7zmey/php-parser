package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Trait struct {
	attributes map[string]interface{}
	position   *node.Position
	TraitName  node.Node
	Stmts      []node.Node
}

func NewTrait(TraitName node.Node, Stmts []node.Node, phpDocComment string) node.Node {
	return &Trait{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		TraitName,
		Stmts,
	}
}

func (n Trait) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Trait) Position() *node.Position {
	return n.position
}

func (n Trait) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Trait) Walk(v node.Visitor) {
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
