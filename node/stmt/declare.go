package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Declare struct {
	position *node.Position
	comments []comment.Comment
	Consts   []node.Node
	Stmt     node.Node
}

func NewDeclare(Consts []node.Node, Stmt node.Node) *Declare {
	return &Declare{
		nil,
		nil,
		Consts,
		Stmt,
	}
}

func (n *Declare) Attributes() map[string]interface{} {
	return nil
}

func (n *Declare) Position() *node.Position {
	return n.position
}

func (n *Declare) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Declare) Comments() []comment.Comment {
	return n.comments
}

func (n *Declare) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Declare) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			nn.Walk(vv)
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
