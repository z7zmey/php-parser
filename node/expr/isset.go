package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Isset struct {
	position  *node.Position
	comments  []comment.Comment
	Variables []node.Node
}

func NewIsset(Variables []node.Node) *Isset {
	return &Isset{
		nil,
		nil,
		Variables,
	}
}

func (n *Isset) Attributes() map[string]interface{} {
	return nil
}

func (n *Isset) Position() *node.Position {
	return n.position
}

func (n *Isset) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Isset) Comments() []comment.Comment {
	return n.comments
}

func (n *Isset) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Isset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variables != nil {
		vv := v.GetChildrenVisitor("Variables")
		for _, nn := range n.Variables {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
