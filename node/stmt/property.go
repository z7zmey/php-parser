package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Property struct {
	position      *node.Position
	comments      *[]comment.Comment
	PhpDocComment string
	Variable      node.Node
	Expr          node.Node
}

func NewProperty(Variable node.Node, Expr node.Node, PhpDocComment string) node.Node {
	return &Property{
		nil,
		nil,
		PhpDocComment,
		Variable,
		Expr,
	}
}
func (n Property) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n Property) Position() *node.Position {
	return n.position
}

func (n Property) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Property) Comments() *[]comment.Comment {
	return n.comments
}

func (n Property) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Property) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
