package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type For struct {
	position *node.Position
	comments *[]comment.Comment
	Init     []node.Node
	Cond     []node.Node
	Loop     []node.Node
	Stmt     node.Node
}

func NewFor(Init []node.Node, Cond []node.Node, Loop []node.Node, Stmt node.Node) node.Node {
	return &For{
		nil,
		nil,
		Init,
		Cond,
		Loop,
		Stmt,
	}
}

func (n For) Attributes() map[string]interface{} {
	return nil
}

func (n For) Position() *node.Position {
	return n.position
}

func (n For) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n For) Comments() *[]comment.Comment {
	return n.comments
}

func (n For) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n For) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Init != nil {
		vv := v.GetChildrenVisitor("Init")
		for _, nn := range n.Init {
			nn.Walk(vv)
		}
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		for _, nn := range n.Cond {
			nn.Walk(vv)
		}
	}

	if n.Loop != nil {
		vv := v.GetChildrenVisitor("Loop")
		for _, nn := range n.Loop {
			nn.Walk(vv)
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
