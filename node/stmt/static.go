package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Static struct {
	position *node.Position
	comments *[]comment.Comment
	Vars     []node.Node
}

func NewStatic(Vars []node.Node) *Static {
	return &Static{
		nil,
		nil,
		Vars,
	}
}

func (n *Static) Attributes() map[string]interface{} {
	return nil
}

func (n *Static) Position() *node.Position {
	return n.position
}

func (n *Static) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Static) Comments() *[]comment.Comment {
	return n.comments
}

func (n *Static) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Static) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		vv := v.GetChildrenVisitor("Vars")
		for _, nn := range n.Vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
