package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type HaltCompiler struct {
	position *node.Position
	comments *[]comment.Comment
}

func NewHaltCompiler() *HaltCompiler {
	return &HaltCompiler{
		nil,
		nil,
	}
}

func (n HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

func (n HaltCompiler) Position() *node.Position {
	return n.position
}

func (n HaltCompiler) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n HaltCompiler) Comments() *[]comment.Comment {
	return n.comments
}

func (n HaltCompiler) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n HaltCompiler) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
