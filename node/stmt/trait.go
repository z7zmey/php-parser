package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Trait struct {
	position      *node.Position
	comments      *[]comment.Comment
	PhpDocComment string
	TraitName     node.Node
	Stmts         []node.Node
}

func NewTrait(TraitName node.Node, Stmts []node.Node, PhpDocComment string) node.Node {
	return &Trait{
		nil,
		nil,
		PhpDocComment,
		TraitName,
		Stmts,
	}
}

func (n Trait) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n Trait) Position() *node.Position {
	return n.position
}

func (n Trait) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Trait) Comments() *[]comment.Comment {
	return n.comments
}

func (n Trait) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
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
