package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	position *node.Position
	comments *[]comment.Comment
	Vars     []node.Node
}

func NewGlobal(Vars []node.Node) *Global {
	return &Global{
		nil,
		nil,
		Vars,
	}
}

func (n *Global) Attributes() map[string]interface{} {
	return nil
}

func (n *Global) Position() *node.Position {
	return n.position
}

func (n *Global) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Global) Comments() *[]comment.Comment {
	return n.comments
}

func (n *Global) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Global) Walk(v node.Visitor) {
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
