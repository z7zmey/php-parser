package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Function struct {
	position      *node.Position
	comments      *[]comment.Comment
	ReturnsRef    bool
	PhpDocComment string
	FunctionName  node.Node
	Params        []node.Node
	ReturnType    node.Node
	Stmts         []node.Node
}

func NewFunction(FunctionName node.Node, ReturnsRef bool, Params []node.Node, ReturnType node.Node, Stmts []node.Node, PhpDocComment string) node.Node {
	return &Function{
		nil,
		nil,
		ReturnsRef,
		PhpDocComment,
		FunctionName,
		Params,
		ReturnType,
		Stmts,
	}
}

func (n Function) Attributes() map[string]interface{} {
	// return n.attributes
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n Function) Position() *node.Position {
	return n.position
}

func (n Function) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Function) Comments() *[]comment.Comment {
	return n.comments
}

func (n Function) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n Function) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.FunctionName != nil {
		vv := v.GetChildrenVisitor("FunctionName")
		n.FunctionName.Walk(vv)
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
