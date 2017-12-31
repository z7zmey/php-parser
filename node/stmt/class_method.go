package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassMethod struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
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
		nil,
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
	return n.attributes
}

func (n ClassMethod) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ClassMethod) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
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
