package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	position *node.Position
	comments []comment.Comment
	VarName  node.Node
}

func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		nil,
		nil,
		VarName,
	}
}

func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

func (n *Variable) Position() *node.Position {
	return n.position
}

func (n *Variable) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Variable) Comments() []comment.Comment {
	return n.comments
}

func (n *Variable) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		vv := v.GetChildrenVisitor("VarName")
		n.VarName.Walk(vv)
	}

	v.LeaveNode(n)
}
