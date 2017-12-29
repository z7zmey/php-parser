package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassMethod struct {
	name       string
	attributes map[string]interface{}
	methodName node.Node
	modifiers  []node.Node
	params     []node.Node
	returnType node.Node
	stmts      []node.Node
}

func NewClassMethod(methodName node.Node, modifiers []node.Node, returnsRef bool, params []node.Node, returnType node.Node, stmts []node.Node) node.Node {
	return ClassMethod{
		"ClassMethod",
		map[string]interface{}{
			"returnsRef": returnsRef,
		},
		methodName,
		modifiers,
		params,
		returnType,
		stmts,
	}
}

func (n ClassMethod) Name() string {
	return "ClassMethod"
}

func (n ClassMethod) Attributes() map[string]interface{} {
	return nil
}

func (n ClassMethod) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.methodName != nil {
		vv := v.GetChildrenVisitor("methodName")
		n.methodName.Walk(vv)
	}

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
