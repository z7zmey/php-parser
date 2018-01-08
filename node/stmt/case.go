package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Case struct {
	position *node.Position
	comments *[]comment.Comment
	Cond     node.Node
	Stmts    []node.Node
}

func NewCase(Cond node.Node, Stmts []node.Node) *Case {
	return &Case{
		nil,
		nil,
		Cond,
		Stmts,
	}
}

func (n Case) Attributes() map[string]interface{} {
	return nil
}

func (n Case) Position() *node.Position {
	return n.position
}

func (n Case) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Case) Comments() *[]comment.Comment {
	return n.comments
}

func (n Case) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Case) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
