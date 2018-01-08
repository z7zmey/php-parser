package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Echo struct {
	position *node.Position
	comments []comment.Comment
	Exprs    []node.Node
}

func NewEcho(Exprs []node.Node) *Echo {
	return &Echo{
		nil,
		nil,
		Exprs,
	}
}

func (n *Echo) Attributes() map[string]interface{} {
	return nil
}

func (n *Echo) Position() *node.Position {
	return n.position
}

func (n *Echo) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Echo) Comments() []comment.Comment {
	return n.comments
}

func (n *Echo) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Echo) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		vv := v.GetChildrenVisitor("Exprs")
		for _, nn := range n.Exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
