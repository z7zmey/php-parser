package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Trait struct {
	attributes map[string]interface{}
	position   *node.Position
	traitName  node.Node
	stmts      []node.Node
}

func NewTrait(traitName node.Node, stmts []node.Node, phpDocComment string) node.Node {
	return &Trait{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		traitName,
		stmts,
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
