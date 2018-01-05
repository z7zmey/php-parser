package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Closure struct {
	position      *node.Position
	comments      *[]comment.Comment
	ReturnsRef    bool
	Static        bool
	PhpDocComment string
	Params        []node.Node
	Uses          []node.Node
	ReturnType    node.Node
	Stmts         []node.Node
}

func NewClosure(Params []node.Node, Uses []node.Node, ReturnType node.Node, Stmts []node.Node, Static bool, ReturnsRef bool, PhpDocComment string) node.Node {
	return &Closure{
		nil,
		nil,
		ReturnsRef,
		Static,
		PhpDocComment,
		Params,
		Uses,
		ReturnType,
		Stmts,
	}
}

func (n Closure) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"Static":        n.Static,
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n Closure) Position() *node.Position {
	return n.position
}

func (n Closure) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Closure) Comments() *[]comment.Comment {
	return n.comments
}

func (n Closure) SetComments(c []comment.Comment) node.Node {
	n.comments = &c
	return n
}

func (n Closure) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Params != nil {
		vv := v.GetChildrenVisitor("Params")
		for _, nn := range n.Params {
			nn.Walk(vv)
		}
	}

	if n.Uses != nil {
		vv := v.GetChildrenVisitor("Uses")
		for _, nn := range n.Uses {
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
