package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Try struct {
	position *node.Position
	comments *[]comment.Comment
	Stmts    []node.Node
	Catches  []node.Node
	Finally  node.Node
}

func NewTry(Stmts []node.Node, Catches []node.Node, Finally node.Node) node.Node {
	return &Try{
		nil,
		nil,
		Stmts,
		Catches,
		Finally,
	}
}

func (n Try) Attributes() map[string]interface{} {
	return nil
}

func (n Try) Position() *node.Position {
	return n.position
}

func (n Try) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Try) Comments() *[]comment.Comment {
	return n.comments
}

func (n Try) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n Try) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	if n.Catches != nil {
		vv := v.GetChildrenVisitor("Catches")
		for _, nn := range n.Catches {
			nn.Walk(vv)
		}
	}

	if n.Finally != nil {
		vv := v.GetChildrenVisitor("Finally")
		n.Finally.Walk(vv)
	}

	v.LeaveNode(n)
}
