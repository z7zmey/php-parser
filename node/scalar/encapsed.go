package scalar

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Encapsed struct {
	position *node.Position
	comments []comment.Comment
	Parts    []node.Node
}

func NewEncapsed(Parts []node.Node) *Encapsed {
	return &Encapsed{
		nil,
		nil,
		Parts,
	}
}

func (n *Encapsed) Attributes() map[string]interface{} {
	return nil
}

func (n *Encapsed) Position() *node.Position {
	return n.position
}

func (n *Encapsed) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Encapsed) Comments() []comment.Comment {
	return n.comments
}

func (n *Encapsed) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Encapsed) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			nn.Walk(vv)
		}
	}
}
