package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type While struct {
	position *node.Position
	comments *[]comment.Comment
	Token    token.Token
	Cond     node.Node
	Stmt     node.Node
}

func NewWhile(Token token.Token, Cond node.Node, Stmt node.Node) node.Node {
	return &While{
		nil,
		nil,
		Token,
		Cond,
		Stmt,
	}
}

func (n While) Attributes() map[string]interface{} {
	return nil
}

func (n While) Position() *node.Position {
	return n.position
}

func (n While) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n While) Comments() *[]comment.Comment {
	return n.comments
}

func (n While) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n While) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
