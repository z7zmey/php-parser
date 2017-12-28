package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ClassMethod) Name() string {
	return "ClassMethod"
}

type ClassMethod struct {
	name        string
	token       token.Token
	modifiers   []node.Node
	isReturnRef bool
	params      []node.Node
	returnType  node.Node
	stmts       []node.Node
}

func NewClassMethod(token token.Token, modifiers []node.Node, isReturnRef bool, params []node.Node, returnType node.Node, stmts []node.Node) node.Node {
	return ClassMethod{
		"ClassMethod",
		token,
		modifiers,
		isReturnRef,
		params,
		returnType,
		stmts,
	}
}

func (n ClassMethod) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
	v.Scalar("isReturnRef", n.isReturnRef)

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.params != nil {
		vv := v.GetChildrenVisitor("params")
		for _, nn := range n.params {
			nn.Walk(vv)
		}
	}

	if n.returnType != nil {
		vv := v.GetChildrenVisitor("returnType")
		n.returnType.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
