package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassMethod struct {
	attributes map[string]interface{}
	position   *node.Position
	MethodName node.Node
	Modifiers  []node.Node
	Params     []node.Node
	ReturnType node.Node
	Stmts      []node.Node
}

func NewClassMethod(MethodName node.Node, Modifiers []node.Node, returnsRef bool, Params []node.Node, ReturnType node.Node, Stmts []node.Node, phpDocComment string) node.Node {
	return &ClassMethod{
		map[string]interface{}{
			"returnsRef":    returnsRef,
			"phpDocComment": phpDocComment,
		},
		nil,
		MethodName,
		Modifiers,
		Params,
		ReturnType,
		Stmts,
	}
}

func (n ClassMethod) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ClassMethod) Position() *node.Position {
	return n.position
}

func (n ClassMethod) SetPosition(p *node.Position) node.Node {
	n.position = p
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
