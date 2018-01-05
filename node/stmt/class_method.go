package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type ClassMethod struct {
	position      *node.Position
	comments      *[]comment.Comment
	ReturnsRef    bool
	PhpDocComment string
	MethodName    node.Node
	Modifiers     []node.Node
	Params        []node.Node
	ReturnType    node.Node
	Stmts         []node.Node
}

func NewClassMethod(MethodName node.Node, Modifiers []node.Node, ReturnsRef bool, Params []node.Node, ReturnType node.Node, Stmts []node.Node, PhpDocComment string) node.Node {
	return &ClassMethod{
		nil,
		nil,
		ReturnsRef,
		PhpDocComment,
		MethodName,
		Modifiers,
		Params,
		ReturnType,
		Stmts,
	}
}

func (n ClassMethod) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n ClassMethod) Position() *node.Position {
	return n.position
}

func (n ClassMethod) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ClassMethod) Comments() *[]comment.Comment {
	return n.comments
}

func (n ClassMethod) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n ClassMethod) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.MethodName != nil {
		vv := v.GetChildrenVisitor("MethodName")
		n.MethodName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			nn.Walk(vv)
		}
	}

	if n.Params != nil {
		vv := v.GetChildrenVisitor("Params")
		for _, nn := range n.Params {
			nn.Walk(vv)
		}
	}

	if n.ReturnType != nil {
		vv := v.GetChildrenVisitor("ReturnType")
		n.ReturnType.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
