package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Finally struct {
	position *node.Position
	comments *[]comment.Comment
	Stmts    []node.Node
}

func NewFinally(Stmts []node.Node) node.Node {
	return &Finally{
		nil,
		nil,
		Stmts,
	}
}

func (n Finally) Attributes() map[string]interface{} {
	return nil
}

func (n Finally) Position() *node.Position {
	return n.position
}

func (n Finally) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Finally) Comments() *[]comment.Comment {
	return n.comments
}

func (n Finally) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Finally) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
